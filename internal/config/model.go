package config

import (
	"os"

	"github.com/slashbaseide/slashbase/internal/utils"
)

type AppConfig struct {
	Version           string
	BuildName         string
	Port              string
	CryptedDataSecret string
	SecurityKey       string
}

func newConfig(buildName, version string) AppConfig {
	return AppConfig{
		Version:           version,
		BuildName:         buildName,
		Port:              DEFAULT_SERVER_PORT,
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
		SecurityKey:       utils.RandString(25),
	}
}
