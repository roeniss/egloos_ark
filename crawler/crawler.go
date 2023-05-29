package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// *******************************
// Setup
//
//

var crawledPosts int
var savedImages int

func InitProgressDisplay() {
	go func() {
		for {
			fmt.Printf("\rCrawled %d pages / Saved %d imagesls | ", crawledPosts, savedImages)
			time.Sleep(time.Millisecond * 100) // Adjust the sleep duration as needed
		}
	}()
}

// *******************************
// Domain
type Post struct {
	PostNum int
	RawHtml string
	BlogId  string
}

// *******************************
// Crawler
//
//

type CrawlingFailedPost struct {
	PostNum int
	Reason  string
	BlogId  string
}

type Crawler struct {
	Collector *colly.Collector
	BlogId    string
}

var ParseAndSavePostRe = regexp.MustCompile(`^/\d+/?$`) // e.g. `/231244`, `/231244/`

var imageCollector *colly.Collector

func CreateCrawler(repo *SqlRepository, blogId string) *Crawler {
	mainCollector := colly.NewCollector(
		colly.UserAgent("EGLOOS_ARK"),
		colly.IgnoreRobotsTxt(),

		colly.CacheDir("./colly_cache_dir"), // GET response cache
		// colly.Debugger(&debug.WebDebugger{}), // too noisy
		// colly.Async(true), // server down happen..
	)

	err := mainCollector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		Delay:       10 * time.Millisecond,
	})
	if err != nil {
		panic(err)
	}
	mainCollector.SetRequestTimeout(10 * time.Second)

	mainCollector.OnError(func(r *colly.Response, err error) {
		postNum := getPostNum(r.Request.URL.Path)

		crawlingFailedPost := &CrawlingFailedPost{
			PostNum: postNum,
			Reason:  err.Error(),
			BlogId:  blogId,
		}
		repo.SaveCrawlingFailedPost(crawlingFailedPost)
		crawledPosts++
	})

	categoriesCrawled := false
	mainCollector.OnHTML(".widget.menu_category .widget_content", func(e *colly.HTMLElement) {
		if categoriesCrawled {
			return
		}
		categoriesCrawled = true

		e.ForEach("li a[href]", func(_ int, e2 *colly.HTMLElement) {
			if e2.Text == "전체" {
				return
			}

			url := e2.Request.AbsoluteURL(e2.Attr("href") + "/page/1")
			err2 := e2.Request.Visit(url)
			if err2 != nil && err2 != colly.ErrAlreadyVisited {
				panic(err2)
			}

		})
	})

	// visit each post (e.g. `/4216814`)
	mainCollector.OnHTML("#titlelist_list a[href]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("href"))
		err2 := e.Request.Visit(url)
		if err2 != nil {
			panic(err2)
		}
	})

	mainCollector.OnHTML("#titlelist_paging", func(e *colly.HTMLElement) {
		if !strings.Contains(e.Request.URL.Path, "/page") {
			return
		}

		e.ForEach("a", func(_ int, e *colly.HTMLElement) {
			url := e.Request.AbsoluteURL(e.Attr("href"))
			err2 := e.Request.Visit(url)
			if err2 != nil && err2 != colly.ErrAlreadyVisited {
				panic(err2)
			}
		})
	})

	// parse each post
	mainCollector.OnHTML("#section_content", func(e *colly.HTMLElement) {
		if !ParseAndSavePostRe.MatchString(e.Request.URL.Path) {
			return
		}
		ParseAndSavePost(e, repo, blogId)
	})

	// setup image collector
	imageCollector = mainCollector.Clone()

	imageCollector.OnResponse(func(r *colly.Response) {
		filename := fmt.Sprintf("./images/%s", url.QueryEscape(r.Request.URL.String()))
		err := r.Save(filename)
		if err != nil {
			fmt.Printf("이미지 저장 실패: %s (err: %v)\n", filename, err)
		} else {
			savedImages++
		}
	})

	return &Crawler{mainCollector, blogId}
}

func ParseAndSavePost(e *colly.HTMLElement, repo *SqlRepository, blogId string) {
	postNum := getPostNum(e.Request.URL.Path)

	if strings.Contains(e.Text, "등록된 포스트가 없습니다.[새글쓰기] 메뉴를 눌러 새로운 포스트를 올리시기 바랍니다.") {
		crawlingFailedPost := &CrawlingFailedPost{
			PostNum: postNum,
			Reason:  "No post",
		}
		repo.SaveCrawlingFailedPost(crawlingFailedPost)

	} else if htmlStr, err2 := e.DOM.Html(); err2 != nil {
		crawlingFailedPost := &CrawlingFailedPost{
			PostNum: postNum,
			Reason:  "Failed to `e.DOM.Html()`",
		}
		repo.SaveCrawlingFailedPost(crawlingFailedPost)

	} else {
		body := e.DOM.Find(".post_content > div > span img") // only images in post body

		SaveImages(body)
		SavePost(postNum, htmlStr, blogId, repo)
	}

	crawledPosts++
}

func SavePost(postNum int, htmlStr string, blogId string, repo *SqlRepository) {
	post := &Post{
		PostNum: postNum,
		RawHtml: MinifyHtml(htmlStr),
		BlogId:  blogId,
	}
	repo.SavePost(post)
}

func SaveImages(body *goquery.Selection) {
	for _, img := range body.Nodes {
		for _, attr := range img.Attr {
			if attr.Key == "src" {
				imageCollector.Visit(attr.Val)
			}
		}
	}
}

func getPostNum(path string) int {
	numStr := regexp.MustCompile(`\d+`).FindString(path)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}

// *******************************
// main
//
//

func (c *Crawler) Crawl() {
	err := c.Collector.Visit("http://" + c.BlogId + ".egloos.com/")
	if err != nil {
		panic(err)
	}

	//c.Collector.Wait() // this is needed for `colly.Async(true)`
}

type CrawlerStarter interface {
	Crawl()
}

func Setup(blogId string) CrawlerStarter {
	// blogId := GetBlogId()

	repo := InitRepository()
	// defer repo.Close()

	InitProgressDisplay()

	return CreateCrawler(repo, blogId)
}
