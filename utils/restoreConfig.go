package utils

import (
	"fmt"
	"os"
	"stu-net/tools"
)

func RestoreConfigFile() {
	configFile := tools.ConfigPath()
	err := os.Remove(configFile)
	if err != nil {
		fmt.Println(err)
	}

	tools.CreateConfigFileIfNotExists(configFile)
}
