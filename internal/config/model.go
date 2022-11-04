package config

import "os"

type AppConfig struct {
	EnvName           string
	Port              string
	AuthTokenSecret   string
	CryptedDataSecret string
}

func newConfig() AppConfig {
	return AppConfig{
		EnvName:           os.Getenv("ENV_NAME"),
		Port:              os.Getenv("PORT"),
		AuthTokenSecret:   os.Getenv("AUTH_TOKEN_SECRET"),
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
	}
}
