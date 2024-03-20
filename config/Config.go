package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CONFIG_PATH      = "../../config/"
	CONFIG_FILE_NAME = "config"
	CONFIG_TYPE      = "env"
)

type envConfig struct {
	DbUrl      string
	MailApiKey string
	MailUrl    string
}

func LoadConfigFile() *envConfig {
	var appConfig = &envConfig{}
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file: ", err)
	}
	if err := viper.Unmarshal(appConfig); err != nil {
		log.Fatal("Error reading config file: ", err)
	}
	return appConfig
}
