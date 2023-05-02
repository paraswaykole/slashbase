package config

import (
	"log"

	"github.com/joho/godotenv"
)

var config AppConfig

func Init(build, envName, version string) {
	if build == BUILD_DESKTOP {
		if envName == ENV_NAME_DEVELOPMENT {
			err := godotenv.Load("development.env")
			if err != nil {
				log.Fatal("Error loading development.env file")
			}
		} else if envName == ENV_NAME_PRODUCTION {
			err := godotenv.Load(GetAppEnvFilePath())
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		}
	} else {
		if envName == ENV_NAME_DEVELOPMENT {
			err := godotenv.Load("development.server.env")
			if err != nil {
				log.Fatal("Error loading development.server.env file")
			}
		}
	}
	config = newConfig(build, envName, version)
}

func IsLive() bool {
	return config.EnvName == ENV_NAME_PRODUCTION
}

func IsDesktop() bool {
	return config.Build == BUILD_DESKTOP
}

func GetConfig() *AppConfig {
	return &config
}
