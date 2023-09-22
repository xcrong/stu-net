package cmd

import (
	"os"

	"stu-net/utils"

	"github.com/spf13/cobra"
)

var interactive bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stu-net",
	Short: "命令行校园网管理工具。",
	Long: `stu-net 是一款可以在命令行中运行的校园网工具。
可以用于无头环境下校园网的登录、登出、状态查询、流量查询。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetUsageTemplate(cmd.UsageString() + "\n" + `Examples:
	stu-net -i # 进入交互模式
	stu-net login \
		-u 20xxxxx \
		-p my_pwd \
		-s # 指定用户名和密码登录，同时将之保存到配置
	stu-net flux # 查询流量
`)
		if interactive {
			utils.InteractiveMode()
			// os.Exit(0)
		} else {
			//cmd.Help()

			err := cmd.Help()
			if err != nil {
				return
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "使用交互模式")
}
