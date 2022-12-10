package config

import (
	"log"
	"os"

	"slashbase.com/backend/internal/utils"
)

type AppConfig struct {
	EnvName           string
	Port              string
	CryptedDataSecret string
}

func newConfig(envName string) AppConfig {
	cryptedDataSecret := os.Getenv("CRYPTED_DATA_SECRET")
	if cryptedDataSecret == "" {
		hex, err := utils.RandomHex(32)
		if err != nil {
			log.Fatal("env CRYPTED_DATA_SECRET not found")
		}
		cryptedDataSecret = hex
	}
	if os.Getenv("ENV_NAME") != "" {
		envName = os.Getenv("ENV_NAME")
	}
	return AppConfig{
		EnvName:           envName,
		Port:              os.Getenv("PORT"),
		CryptedDataSecret: cryptedDataSecret,
	}
}
