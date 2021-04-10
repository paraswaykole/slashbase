package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

// GetConfig from viper
func GetConfig() *viper.Viper {
	return config
}

// IsLive gives if the current environment is production
func IsLive() bool {
	return config.GetString("name") == "prod"
}

// IsStage gives if the current environment is stage
func IsStage() bool {
	return config.GetString("name") == "stage"
}
