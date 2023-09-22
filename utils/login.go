package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"stu-net/tools"
)

func Login(account *Account) (msg *tools.LogResult, cookie string, err error) {
	// 创建一个 HTTP 客户端，并忽略SSL 错误
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := url.Values{
		"opr":         {"pwdLogin"},
		"userName":    {account.Username},
		"pwd":         {account.Password},
		"ipv4or6":     {""},
		"rememberPwd": {"0"},
	}

	resp, err := client.Post(LoginUrl, "application/x-www-form-urlencoded", strings.NewReader(payload.Encode()))
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}

	bodyString := string(body)
	msg, err = tools.ParseLogInfo(bodyString)
	if err != nil {
		// fmt.Println(err)
		return nil, "", err
	}

	cookieSet := resp.Cookies()
	for _, c := range cookieSet {
		if c.Name == "AUTHSESSID" {
			cookie = c.Value
		}
	}

	return msg, cookie, nil
}
