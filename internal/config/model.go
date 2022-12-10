package config

import "os"

type AppConfig struct {
	EnvName           string
	Port              string
	CryptedDataSecret string
}

func newConfig() AppConfig {
	return AppConfig{
		EnvName:           os.Getenv("ENV_NAME"),
		Port:              os.Getenv("PORT"),
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
	}
}
