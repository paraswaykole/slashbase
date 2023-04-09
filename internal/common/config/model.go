package config

import (
	"os"
)

type AppConfig struct {
	Version           string
	Build             string
	EnvName           string
	Port              string
	CryptedDataSecret string
}

func newConfig(build, envName, version string) AppConfig {
	return AppConfig{
		Version:           version,
		Build:             build,
		EnvName:           envName,
		Port:              DEFAULT_SERVER_PORT,
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
	}
}
