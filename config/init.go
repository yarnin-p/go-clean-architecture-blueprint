package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config")

	return viper.ReadInConfig()
}
