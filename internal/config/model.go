package config

import (
	"log"
	"os"

	"github.com/slashbaseide/slashbase/internal/utils"
)

type AppConfig struct {
	BuildName         string
	Port              string
	CryptedDataSecret string
}

func newConfig(buildName string) AppConfig {
	cryptedDataSecret := os.Getenv("CRYPTED_DATA_SECRET")
	if cryptedDataSecret == "" {
		hex, err := utils.RandomHex(32)
		if err != nil {
			log.Fatal("env CRYPTED_DATA_SECRET not found")
		}
		cryptedDataSecret = hex
	}
	return AppConfig{
		BuildName:         buildName,
		Port:              os.Getenv("PORT"),
		CryptedDataSecret: cryptedDataSecret,
	}
}
