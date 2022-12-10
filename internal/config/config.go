package config

import (
	"log"

	"github.com/joho/godotenv"
)

var config AppConfig

func Init(env string) {
	if env == ENV_DEVELOPMENT {
		err := godotenv.Load("development.env")
		if err != nil {
			log.Fatal("Error loading development.env file")
		}
	}
	config = newConfig()
}

func IsLive() bool {
	return config.EnvName == ENV_PRODUCTION
}

func GetConfig() *AppConfig {
	return &config
}

func GetServerPort() string {
	if config.Port == "" {
		return DEFAULT_SERVER_PORT
	}
	return config.Port
}
