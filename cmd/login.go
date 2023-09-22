package cmd

import (
	"fmt"
	"stu-net/utils"

	"github.com/spf13/cobra"
)

var uname string
var pwd string
var toSave bool

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "通过账号和密码登录到校园网",
	Long: `通过帐号和密码登录到校园网
这会比在浏览器登录更快，因为程序直接向服务器
发送请求，不需要请求额外资源。`,
	Run: func(cmd *cobra.Command, args []string) {
		checkResult, err := utils.Check()
		if err != nil {
			fmt.Println(err)
			return
		}
		if checkResult.Online == 1 {
			fmt.Printf("用户 %s 已登录...\n", checkResult.Username)
			return
		}
		if uname == "" || pwd == "" {
			fmt.Println("未指定用户名或密码...")
			fmt.Println("尝试读取配置文件...")
			account, err := utils.ReadAccount()

			if err != nil {
				fmt.Println(err)
				fmt.Println("读取配置文件失败...")
				fmt.Println("请指定用户名和密码...\nTip： 加上参数 -s 保存到配置文件\nExample: stu-login login -u 校园网用户名 -p 校园网密码 -s")
				return
			}
			uname = account.Username
			pwd = account.Password

			loginResult, cookie, err := utils.Login(account)
			if err != nil {
				fmt.Println(err)
				return
			}

			utils.ParseLoginResult(loginResult, cookie)
			return
		}
		account := &utils.Account{
			Username: uname,
			Password: pwd,
		}
		if toSave {
			_, err = utils.StoreAccount(account)
			if err != nil {
				fmt.Println(err)
				return 
			}
		}

		loginResult, cookie, err := utils.Login(account)
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.ParseLoginResult(loginResult, cookie)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&uname, "username", "u", "", "你的校园网用户名  (如果密码设置则为必须)")
	loginCmd.Flags().StringVarP(&pwd, "password", "p", "", "你的校园网密码  (如果用户名设置则为必须)")
	loginCmd.MarkFlagsRequiredTogether("username", "password")
	loginCmd.Flags().BoolVarP(&toSave, "save", "s", false, "保存用户名和密码到配置文件")

}
