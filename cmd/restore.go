package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"stu-net/tools"
	"stu-net/utils"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "重置配罝文件",
	Long: `
重置配罝文件。`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.RestoreConfigFile()
		fmt.Println(tools.ConfigPath(), "己经重置...")
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
