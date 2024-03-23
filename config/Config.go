package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CONFIG_PATH      = "config/"
	CONFIG_FILE_NAME = "config"
	CONFIG_TYPE      = "env"
)

type EnvConfig struct {
	DbUrl        string
	MAIL_API_KEY string
	MAIL_API_URL string
}

func LoadConfigFile() *EnvConfig {
	var appConfig *EnvConfig
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_FILE_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error loading config file: ", err)
	}
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatal("Error reading config file: ", err)
	}
	return appConfig
}
