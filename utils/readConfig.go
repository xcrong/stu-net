package utils

import (
	"fmt"
	"stu-net/tools"
	"time"

	"github.com/go-ini/ini"
)

func ReadAccount() (account *Account, err error) {
	cfg, err := ini.Load(tools.ConfigPath())
	if err != nil {
		return nil, err
	}

	AccountSection := cfg.Section("Account")
	enctyptedUname := AccountSection.Key("uname").String()
	enctyptedPwd := AccountSection.Key("pwd").String()
	key := AccountSection.Key("key").String()
	verify := AccountSection.Key("verify").String()

	if verify != tools.GenMD5(enctyptedUname+"."+enctyptedPwd+"."+key) {
		return nil, fmt.Errorf("配置文件 %s 无效", tools.ConfigPath())
	}

	uname, err := tools.Decrypt(enctyptedUname, tools.SortKey(key))
	if err != nil || uname == "" {
		return nil, err
	}

	pwd, err := tools.Decrypt(enctyptedPwd, tools.SortKey(key))
	if err != nil || pwd == "" {
		return nil, err
	}

	account = &Account{
		Username: uname,
		Password: pwd,
	}

	return account, nil
}

func ReadCookie() (cookie string, err error) {
	cfg, err := ini.Load(tools.ConfigPath())
	if err != nil {
		return "", err
	}

	CacheSection := cfg.Section("Cache")
	timeStr := CacheSection.Key("time").String()
	storeTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return "", err
	}

	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	if storeTime.Before(yesterday) {
		return "", fmt.Errorf("cookie 已经过期，请重新获取。")
	}

	cookie = CacheSection.Key("cookie").String()
	if cookie == "" {
		return "", fmt.Errorf("cookie 为空，请重新获取。")
	}
	return cookie, nil
}
