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

func Logout() (*tools.LogResult, error) {
	// 创建一个 HTTP 客户端，并忽略SSL 错误
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := url.Values{

		"opr":     {"logout"},
		"ipv4or6": {""},
	}

	resp, err := client.Post(LogoutUrl, "application/x-www-form-urlencoded", strings.NewReader(payload.Encode()))
	if err != nil {
		// fmt.Println(err)
		return nil, err
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
		return nil, err
	}

	logRusult, err := tools.ParseLogInfo(string(body))
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}

	return logRusult, nil
}
