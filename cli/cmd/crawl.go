package cmd

import (
    `github.com/roeniss/EgloosArk/crawler`
    `github.com/spf13/cobra`
)

var CrawlCmd = &cobra.Command{
    Use:   "crawl",
    Short: "특정 이글루스 블로그의 글을 전부 크롤링합니다.",
    Long: `이글루스 리소스를 절약하기 위해 상대적으로 느린 속도로 진행됩니다. 
(일반적으로 1000개 게시물이라면 5분 정도 소요됩니다)
크롤링한 데이터베이스는 현재 디렉토리에 'egloosark_$blogid.db' 
파일로 저장되며, sqlite3 포맷입니다.
비공개 글은 크롤링 되지 않습니다.`,
    Run: func(cmd *cobra.Command, args []string) {
        crawler.Setup(BlogId).Crawl()
    },
}

var BlogId string

func init() {
    CrawlCmd.Flags().StringVarP(&BlogId, "blogid", "b", "", "크롤링할 블로그의 아이디를 지정합니다.")
    CrawlCmd.MarkFlagRequired("blogid")

}
