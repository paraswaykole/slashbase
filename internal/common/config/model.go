package config

import (
	"os"
)

type AppConfig struct {
	Version           string
	Build             string
	EnvName           string
	Port              string
	AuthTokenSecret   string
	CryptedDataSecret string
	AppDB             AppDBConfig
	RootUser          RootUser
}

type AppDBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type RootUser struct {
	Email    string
	Password string
}

func newConfig(build, envName, version string) AppConfig {
	return AppConfig{
		Version:           version,
		Build:             build,
		EnvName:           envName,
		Port:              DEFAULT_SERVER_PORT,
		AuthTokenSecret:   os.Getenv("AUTH_TOKEN_SECRET"),
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),
		AppDB: AppDBConfig{
			Host: os.Getenv("APP_DB_HOST"),
			Port: os.Getenv("APP_DB_PORT"),
			User: os.Getenv("APP_DB_USER"),
			Pass: os.Getenv("APP_DB_PASS"),
			Name: os.Getenv("APP_DB_NAME"),
		},
		RootUser: RootUser{
			Email:    os.Getenv("ROOT_USER_EMAIL"),
			Password: os.Getenv("ROOT_USER_PASSWORD"),
		},
	}
}
