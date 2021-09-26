package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")

	viper.SetConfigName("cfg")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Could not read config file: %v", err)
	}
}
