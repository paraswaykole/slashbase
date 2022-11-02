package config

import "os"

type Config struct {
	EnvName string
	Port    string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	AuthTokenSecret   string
	CryptedDataSecret string

	RootUserEmail    string
	RootUserPassword string

	TelemetryID string
}

func newConfig() Config {
	return Config{
		EnvName: os.Getenv("ENV_NAME"),
		Port:    os.Getenv("PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),

		AuthTokenSecret:   os.Getenv("AUTH_TOKEN_SECRET"),
		CryptedDataSecret: os.Getenv("CRYPTED_DATA_SECRET"),

		RootUserEmail:    os.Getenv("ROOT_USER_EMAIL"),
		RootUserPassword: os.Getenv("ROOT_USER_PASSWORD"),

		TelemetryID: os.Getenv("TELEMETRY_ID"),
	}
}
