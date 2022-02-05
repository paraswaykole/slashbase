package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	if env != ENV_PRODUCTION {
		config.AddConfigPath("../src/config/")
		config.AddConfigPath("src/config/")
	} else {
		config.AddConfigPath(".")
		config.AddConfigPath("./config/")
	}
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
}

const (
	VERSION = "v1.0.2-beta"

	PAGINATION_COUNT = 20

	SESSION_COOKIE_NAME    = "session"
	SESSION_COOKIE_MAX_AGE = 30 * 24 * 60 * 60 * 1000

	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"
	ENV_LOCAL       = "local"
)

func IsLive() bool {
	return config.GetString("name") == ENV_PRODUCTION
}

func IsDevelopment() bool {
	return config.GetString("name") == ENV_DEVELOPMENT
}

func GetServerPort() string {
	return config.GetString("server.port")
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     config.GetString("database.host"),
		User:     config.GetString("database.user"),
		Port:     config.GetString("database.port"),
		Password: config.GetString("database.password"),
		Database: config.GetString("database.database"),
	}
}

func GetAuthTokenSecret() []byte {
	tokensecret := config.GetString("secret.auth_token_secret")
	return []byte(tokensecret)
}

func GetCryptedDataSecretKey() string {
	return config.GetString("secret.crypted_data_secret_key")
}

func GetApiHost() string {
	return config.GetString("constants.api_host")
}

func GetAppHost() string {
	return config.GetString("constants.app_host")
}

func GetTelemetryId() string {
	return config.GetString("telemetry.id")
}

func GetRootUser() *RootUserConfig {
	return &RootUserConfig{
		Email:    config.GetString("root_user.email"),
		Password: config.GetString("root_user.password"),
	}
}
