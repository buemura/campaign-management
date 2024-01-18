package config

import (
	"github.com/spf13/viper"
)

var (
	PORT string
)

func LoadConfigs() {
	loadEnvVariables()
}

func loadEnvVariables() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load environment variables")
	}

	PORT = viper.GetString("PORT")
}
