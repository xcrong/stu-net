package utils

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {

	t.Run("检查已登录时的返回值是否正常", func(t *testing.T) {

		checkResult, err := Check()
		if err != nil {
			t.Error(err)
		}

		if checkResult.Online == 1 {
			fmt.Printf("username: %s\n", checkResult.Username)
		}
	})

	t.Run("检查未登录时的返回值是否正常", func(t *testing.T) {

		Logout()
		checkResult, err := Check()

		if err != nil {
			t.Error(err)
		}

		if checkResult.Online == 1 {
			t.Error("当前应处于未登录状态。")
		}
	})

}
