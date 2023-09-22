package tools

import (
	"os"
	"path/filepath"
)

func ConfigPath() (configPath string) {

	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	configDir := filepath.Join(userCacheDir, "stu-net")

	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
		return filepath.Join(configDir, "config.ini")
	}

	return filepath.Join(configDir, "config.ini")
}
