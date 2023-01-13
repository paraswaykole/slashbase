package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Version           string
	BuildName         string
	Port              string
	CryptedDataSecret string
	Cli               bool
}

func newConfig(buildName, version string) AppConfig {

	cli, err := strconv.ParseBool(os.Getenv("cli"))
	if err != nil {
		cli = true
	}

	return AppConfig{
		Version:           version,
		BuildName:         buildName,
		Port:              DEFAULT_SERVER_PORT,
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
		Cli:               cli,
	}
}
