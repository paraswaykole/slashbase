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
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

// IsLive gives if the current environment is production
func IsLive() bool {
	return config.GetString("name") == "prod"
}

// IsStage gives if the current environment is stage
func IsStage() bool {
	return config.GetString("name") == "stage"
}

func GetServerPort() string {
	return config.GetString("server.port")
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     config.GetString("database.host"),
		User:     config.GetString("database.user"),
		Password: config.GetString("database.password"),
		Database: config.GetString("database.database"),
	}
}

func GetMagicLinkTokenSecret() []byte {
	tokensecret := config.GetString("secret.magic_link_token_secret")
	return []byte(tokensecret)
}

func GetAuthTokenSecret() []byte {
	tokensecret := config.GetString("secret.auth_token_secret")
	return []byte(tokensecret)
}

func GetAppHost() string {
	return config.GetString("constants.app_host")
}
