package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "egloos_ark",
	Short: "이글루스 블로그의 글을 백업하는 프로그램",
	Long: `이글루스 블로그 팀이 서비스를 종료하면서, 
혹시라도 블로그의 글을 추출해주는 솔루션을
만들지 않을까봐 준비하게 되었습니다. 
https://github.com/roeniss/EgloosArr/issues 에서 문제를 제보해주세요.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.AddCommand(CrawlCmd)

	if err := rootCmd.Execute(); err != nil {
	}
}
