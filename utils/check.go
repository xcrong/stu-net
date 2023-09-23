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

// Check 检查是否在线
func Check() (*tools.CheckResult, error) {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	payload := url.Values{
		// opr: online_check
		"opr": {"online_check"},
	}

	resp, err := client.Post(CheckUrl, "application/x-www-form-urlencoded", strings.NewReader(payload.Encode()))
	if err != nil {
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
		return nil, err
	}

	jsonStr := string(body)
	// fmt.Println(jsonStr)

	checkResult, err := tools.ParseCheckResult(jsonStr)
	if err != nil {
		return nil, err
	}
	return checkResult, nil

}
