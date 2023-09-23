package tools

import (
	"os"
)

const configTemp = `[Account]
uname = 
pwd =
key=
verify=

[Cache]
time=
cookie=
`

func CreateConfigFileIfNotExists(filePath string) {

	if !isFileExists(filePath) {
		file, err := createFile(filePath)
		if err != nil {
			panic(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)

		_, err = file.WriteString(configTemp)
		if err != nil {
			panic(err)
		}

		err = file.Sync() // 确保数据被刷新到磁盘
		if err != nil {
			panic(err)
		}
	}

}

func isFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createFile(filePath string) (*os.File, error) {

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
