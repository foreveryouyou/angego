package main

import (
	"fmt"
	"github.com/foreveryouyou/angego/cmd/commands/web"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "angego"}
	rootCmd.Long = "小工具,生成自己用的web项目框架(gin)基础代码。"

	var cmdWeb = &cobra.Command{
		Use:   "web",
		Short: "创建基础web项目",
		Run: func(cmd *cobra.Command, args []string) {
			web.Web()
		},
	}
	var cmdApi = &cobra.Command{
		Use:   "api",
		Short: "创建基础api项目(暂未实现)",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("暂未实现")
		},
	}

	rootCmd.AddCommand(cmdWeb)
	rootCmd.AddCommand(cmdApi)
	rootCmd.Execute()
}
