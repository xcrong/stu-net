package cmd

import (
	"fmt"
	"stu-net/utils"

	"github.com/spf13/cobra"
)

// fluxCmd represents the flux command
var fluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "查询剩余流量",
	Long:  `查询当前流量总额、当前剩余流量。`,
	Run: func(cmd *cobra.Command, args []string) {
		checkResult, err := utils.Check()
		if err != nil {
			fmt.Println(err)
			return
		}
		if checkResult.Online == 0 {
			fmt.Println("请先登录。")
			return
		}
		cookie, err := utils.ReadCookie()
		if err != nil {
			fmt.Println(err)
			return
		}
		fluxResult, err := utils.Flux(cookie)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("用户名: %s\n", fluxResult.Username)
		fmt.Printf("流量总量: %s\n", fluxResult.Total)
		fmt.Printf("已用流量: %s\n", fluxResult.Usage)
		fmt.Printf("过期时间: %s\n", fluxResult.Overdue)
		fmt.Printf("状态: %s\n", fluxResult.Status)
	},
}

func init() {
	rootCmd.AddCommand(fluxCmd)
}
