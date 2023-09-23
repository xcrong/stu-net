package utils

import (
	"testing"

	"stu-net/tools"
)

func TestAccountStore(t *testing.T) {

	tools.CreateConfigFileIfNotExists(tools.ConfigPath())

	account := Account{
		Username: "djfladkj",
		Password: "dflajkg",
	}

	_, err := StoreAccount(&account)
	if err != nil {
		t.Error(err)
	}

	readAccount, err := ReadAccount()
	if err != nil {
		t.Error(err)
	}

	if readAccount.Username != account.Username || readAccount.Password != account.Password {
		t.Error("账号密码错误")
	}
}

func TestStoreCookie(t *testing.T) {
	tools.CreateConfigFileIfNotExists(tools.ConfigPath())

	cookie := "djfladkj"
	_, err := StoreCookie(cookie)
	if err != nil {
		t.Error(err)
	}

	readCookie, err := ReadCookie()
	if err != nil {
		t.Error(err)
	}

	if readCookie != cookie {
		t.Error("Cookie错误")
	}
}
