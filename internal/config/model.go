package config

import (
	"os"
)

type AppConfig struct {
	Version           string
	BuildName         string
	Port              string
	CryptedDataSecret string
}

func newConfig(buildName, version string) AppConfig {
	return AppConfig{
		Version:           version,
		BuildName:         buildName,
		Port:              DEFAULT_SERVER_PORT,
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
	}
}
