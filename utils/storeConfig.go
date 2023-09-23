package utils

import (
	"time"

	"stu-net/tools"

	"github.com/go-ini/ini"
	"github.com/twharmon/gouid"
)

func StoreAccount(account *Account) (bool, error) {
	uname := account.Username
	pwd := account.Password
	key := gouid.String(32, gouid.MixedCaseAlphaNum)

	encrypedUname, err := tools.Encrypt(uname, tools.SortKey(key))
	if err != nil {
		return false, err
	}

	encryptedPwd, err := tools.Encrypt(pwd, tools.SortKey(key))
	if err != nil {
		return false, err
	}

	verify := tools.GenMD5(encrypedUname + "." + encryptedPwd + "." + key)

	cfg, err := ini.Load(tools.ConfigPath())
	if err != nil {
		return false, err
	}

	accountSection := cfg.Section("Account")
	accountSection.Key("uname").SetValue(encrypedUname)
	accountSection.Key("pwd").SetValue(encryptedPwd)
	accountSection.Key("key").SetValue(key)
	accountSection.Key("verify").SetValue(verify)

	err = cfg.SaveTo(tools.ConfigPath())
	if err != nil {
		return false, err
	}
	return true, nil

}

func StoreCookie(cookie string) (bool, error) {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")

	cfg, err := ini.Load(tools.ConfigPath())
	if err != nil {
		return false, err
	}

	cacheSection := cfg.Section("Cache")
	cacheSection.Key("time").SetValue(timeStr)
	cacheSection.Key("cookie").SetValue(cookie)

	err = cfg.SaveTo(tools.ConfigPath())
	if err != nil {
		return false, err
	}

	return true, nil
}
