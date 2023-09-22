package cmd

import (
	"fmt"

	"stu-net/utils"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出校园网",
	Long: `如果你不想让电脑继续访问网络或想要更换其他账号，
请使用此命令登出`,
	Run: func(cmd *cobra.Command, args []string) {
		logoutResult, err := utils.Logout()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(logoutResult.Message)
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
