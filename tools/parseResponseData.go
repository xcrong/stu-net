package tools

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/valyala/fastjson"
)

type LogResult struct {
	Success  bool
	Message  string
	Action   string
	Pop      int
	UserName string
	Location string
}

type FluxResult struct {
	Total    string
	Usage    string
	Overdue  string
	Status   string
	Username string
}

type CheckResult struct {
	Success  bool
	Online   int
	Username string
}

func ParseLogInfo(data string) (*LogResult, error) {

	jsonString := strings.ReplaceAll(data, "'", "\"")

	// 解析JSON字符串
	var p fastjson.Parser
	v, err := p.Parse(jsonString)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return nil, err
	}

	// 使用fastjson来获取各个字段的值
	success := v.GetBool("success")
	msg := string(v.GetStringBytes("msg"))
	action := string(v.GetStringBytes("action"))
	pop := v.GetInt("pop")
	userName := string(v.GetStringBytes("userName"))
	location := string(v.GetStringBytes("location"))

	result := LogResult{
		Success:  success,
		Message:  msg,
		Action:   action,
		Pop:      pop,
		UserName: userName,
		Location: location,
	}

	return &result, nil

}

func ParseFluxInfo(htmlData string) (result *FluxResult, err error) {
	// 使用正则表达式来提取信息
	re := regexp.MustCompile(`<td>([^<]+)</td>\s*<td>([^<]+)</td>`)
	matches := re.FindAllStringSubmatch(htmlData, -1)

	// 创建一个map来存储提取的信息
	data := make(map[string]string)

	var keys = []string{"Username", "Total", "Usage", "Overdue", "Status"}
	// 将提取的信息存储到map中
	for index, match := range matches {
		key := keys[index]
		value := match[2]
		data[key] = value
	}

	fluxResult := &FluxResult{
		Total:    data["Total"],
		Usage:    data["Usage"],
		Overdue:  data["Overdue"],
		Status:   data["Status"],
		Username: data["Username"],
	}

	return fluxResult, nil

}

func ParseCheckResult(jsonData string) (*CheckResult, error) {
	// {"success":true, "online":0}

	var p fastjson.Parser
	v, err := p.Parse(jsonData)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return nil, err
	}

	success := v.GetBool("success")
	online := v.GetInt("online")
	username := v.GetStringBytes("username")

	result := CheckResult{
		Success:  success,
		Online:   online,
		Username: string(username),
	}

	return &result, nil
}
