/*
Copyright © 2023
*/
package main

import (
	"stu-net/cmd"
	"stu-net/tools"
)

func init() {
	configPath := tools.ConfigPath()
	tools.CreateConfigFileIfNotExists(configPath)
}

func main() {
	cmd.Execute()
}
