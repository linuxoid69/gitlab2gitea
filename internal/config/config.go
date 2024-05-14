/*
Copyright © 2024 Rustam Tagaev linuxoid69@gmail.com
*/
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	CONFIG_FILE_NAME = ".gitlab2gitea.yml"
)
func NewConfig() string {
	configFile := fmt.Sprintf("%s/%s", os.Getenv("HOME"), CONFIG_FILE_NAME)

	if _, err := os.Stat(configFile); err != nil {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configFile)
		viper.Set("gitlab.url", "")
		viper.Set("gitlab.token", "")
		viper.Set("gitea.url", "")
		viper.Set("gitea.token", "")
		viper.WriteConfig()

		return fmt.Sprintf("Config file has been created at `%s`", configFile)
	}

	return fmt.Sprintf("Configuration already exists `%s`", configFile) 
}

func CheckConfigFileExists() {
	configFile := fmt.Sprintf("%s/%s", os.Getenv("HOME"), CONFIG_FILE_NAME)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file: ", err)
		fmt.Println("Run `./gitlab2gitea init` for create config file")
		os.Exit(1)
	}
}
