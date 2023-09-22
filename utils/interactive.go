package utils

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"stu-net/tools"
	"time"
)

var option = `>>>>>>>>>>>>> STU-NET <<<<<<<<<<<<<
请通过数字选择要执行的操作：
	0. 重置配置文件
	1. 登录
	2. 退出登录
	3. 流量查询
	4. 状态查询（查看当前是否已登录）
	5. 退出程序`

func InteractiveMode() {
	for {
		fmt.Println(option)
		var choice int
		fmt.Print("请输入> ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
//			return
			continue
		}
		switch choice {
		case 0:
			RestoreConfigFile()
			fmt.Println(tools.ConfigPath(), "己经重置...")
			time.Sleep(time.Millisecond * 200)
		case 1: // Login
			InteractiveLogin()
			fmt.Print("请按回车（Enter）键继续...")
			_, err := fmt.Scanln()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 2: // Logout
			logoutResult, err := Logout()
			if err != nil {
				fmt.Println(err)

			}
			fmt.Println(logoutResult.Message)
			fmt.Print("请按回车（Enter）键继续...")
			_, err = fmt.Scanln()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 3: // Flux
			checkResult, err := Check()
			if err != nil {
				fmt.Println(err)
				return
			}
			if checkResult.Online == 0 {
				fmt.Println("尚未登录")
				fmt.Print("请按回车（Enter）键继续...")
				_, err := fmt.Scanln()
				if err != nil {
					return 
				}
				continue
			}
			cookie, err := ReadCookie()
			if err != nil {
				fmt.Println(err)
				return
			}
			fluxResult, err := Flux(cookie)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("用户名: %s\n", fluxResult.Username)
			fmt.Printf("流量总量: %s\n", fluxResult.Total)
			fmt.Printf("已用流量: %s\n", fluxResult.Usage)
			fmt.Printf("过期时间: %s\n", fluxResult.Overdue)
			fmt.Printf("状态: %s\n", fluxResult.Status)
			fmt.Print("请按回车（Enter）键继续...")
			_, err = fmt.Scanln()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 4: // Check
			checkResult, err := Check()
			if err != nil {
				fmt.Println("发生如下错误，请确认连接到校园网后重试：")
				fmt.Println(err)
				//return
			} else {
				if checkResult.Online == 1 {

					fmt.Printf("已经登录 \n当前用户名：%s\n", checkResult.Username)
				} else {
					fmt.Println("当前未登录")
				}

			}
			fmt.Print("请按回车（Enter）键继续...")
			_, err = fmt.Scanln()
			if err != nil {
				return
			}
		case 5: // Exit
			fmt.Println("goodbye~")
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重新输入")
			fmt.Print("请按回车（Enter）键继续...")
			_, err := fmt.Scanln()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

	}
}

func InteractiveLogin() {
	fmt.Println("尝试读取配置文件...")
	account, err := ReadAccount()
	if err != nil {
		fmt.Println("读取配置文件失败...")
		fmt.Println(err)
		fmt.Println("请手动输入用户名和密码...")
		loginWithInput()
		return
	}

	loginResult, cookie, err := Login(account)
	if err != nil {
		fmt.Println(err)
		fmt.Println("请确保己连接校园网...")
		return
	}

	ParseLoginResult(loginResult, cookie)

}

func loginWithInput() {
	var uname string
	var pwd string
	var toSave string

	fmt.Print("用户名\n>")
	_, err := fmt.Scanln(&uname)
	if err != nil {
		return 
	}

	fmt.Print("密码（输入时不会显示）\n>")
	pwdBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	pwd = string(pwdBytes)

	account := &Account{
		Username: uname,
		Password: pwd,
	}

	for {
		fmt.Println("是否保存用户名和密码到配置文件?(Y/n)\n(回车默认Y)>")
		_, err := fmt.Scanln(&toSave)
		if err != nil {
			return 
		}

		toSave = strings.ToLower(toSave)

		switch toSave {
		case "y", "yes", "":
			_, err := StoreAccount(account)
			if err != nil {
				return 
			}
		case "n", "no":
			fmt.Println("未保存用户名和密码...")
		default:
			fmt.Println("输入错误，请重新输入。Y=yes, n=no")
			continue
		}
		break
	}

	loginResult, cookie, err := Login(account)
	if err != nil {
		fmt.Println(err)
		return
	}

	ParseLoginResult(loginResult, cookie)
}

func ParseLoginResult(loginResult *tools.LogResult, cookie string) {
	if !loginResult.Success {
		fmt.Printf("用户名: %s 登录失败\n", loginResult.UserName)
		fmt.Println(loginResult.Message)
//		fmt.Println(loginResult)
		if loginResult.Message == "用户名或密码错误" {
			RestoreConfigFile()
		}

	} else {
		fluxResult, err := Flux(cookie)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf(` > 登录成功 < 
用户名 ： %s
流量总量 ： %s
已用流量 ： %s

have a good time ~
`, loginResult.UserName, fluxResult.Total, fluxResult.Usage)
		_, err = StoreCookie(cookie)
		if err != nil {
			return
		}
		return
	}
}
