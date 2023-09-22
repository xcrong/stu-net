package utils

import (
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	account := Account{
		Username: "djfladkj",
		Password: "dflajkg",
	}

	msg, cookie, err := Login(&account)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(msg)
	fmt.Println(cookie)
}
