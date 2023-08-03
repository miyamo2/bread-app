/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/miyamo2theppl/bread-app/bread-app-cli/infrastructure/contentful"
	"github.com/miyamo2theppl/bread-app/bread-app-cli/infrastructure/dao"
	"github.com/miyamo2theppl/bread-app/bread-app-cli/services"
	"github.com/miyamo2theppl/bread-app/bread-app-common/core"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "contentful API を使用し、パンの情報を取得し、DBに保存します。",
	Long: `contentful API を使用し、パンの情報を取得し、DBに保存します。
	Example: bread-app-cli fetch bread1 bread2 bread3`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := cfg.Contentful.Token
		bdao := dao.NewBreadDao(core.GetContext(), core.GetClient())
		cclient := contentful.NewClient(token, new(http.Client))
		if len(args) <= 0 {
			cobra.CheckErr(fmt.Errorf("取得するパンは1つ以上指定してください。"))
		}
		svc := services.NewFetchBreadService(cclient, bdao)
		err := svc.Fetch(args)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
