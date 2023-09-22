package cmd

import (
	"fmt"
	"os"

	"stu-net/utils"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "查看当前是否登录",
	Long: `查看当前的登录状态。 
如果已经登录，则打印 true 和 当前用户名
如果未登录，则打印 false`,
	Run: func(cmd *cobra.Command, args []string) {
		checkResult, err := utils.Check()
		if err != nil {
			fmt.Println("发生如下错误，请确认连接到校园网后重试：")
			fmt.Printf("ERROR: %s", err)
			os.Exit(1)
		}
		if checkResult.Online == 1 {

			fmt.Printf("已经登录 \n当前用户名：%s\n", checkResult.Username)
			return
		}
		fmt.Println("当前未登录")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
