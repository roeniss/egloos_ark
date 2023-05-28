package crawler

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

// *******************************
// DB
//
//

type SqlRepository struct {
	db *sql.DB
}

func SetupDbTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS posts (post_num INTEGER, raw_html TEXT, blog_id TEXT, PRIMARY KEY (post_num, blog_id))")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS crawling_failed_posts (post_num INTEGER, reason TEXT, blog_id TEXT, PRIMARY KEY (post_num, blog_id))")
	if err != nil {
		panic(err)
	}
}

func InitRepository() *SqlRepository {
	db, err := sql.Open("sqlite3", "./egloos_ark.db")
	if err != nil {
		panic(err)
	}

	SetupDbTable(db)

	return &SqlRepository{db}
}

func (repo *SqlRepository) Close() {
	err := repo.db.Close()
	if err != nil {
		panic(err)
	}
}

func (repo *SqlRepository) SavePost(post *Post) {
	_, err := repo.db.Exec("INSERT OR IGNORE INTO posts (post_num, raw_html, blog_id) VALUES (?, ?, ?)", post.PostNum, post.RawHtml, post.BlogId)
	if err != nil {
		panic(err)
	}
}

func (repo *SqlRepository) SaveCrawlingFailedPost(post *CrawlingFailedPost) {
	_, err := repo.db.Exec("INSERT OR IGNORE INTO crawling_failed_posts (post_num, reason, blog_id) VALUES (?, ?, ?)", post.PostNum, post.Reason, post.BlogId)
	if err != nil {
		panic(err)
	}
}

// *******************************
// minify
//
//

func MinifyHtml(htmlStr string) string {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	s, err := m.String("text/html", htmlStr)
	if err != nil {
		panic(err)
	}
	return s
}
