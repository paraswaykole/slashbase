package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	VERSION = "v1.0.0-beta"

	PAGINATION_COUNT = 20

	SESSION_COOKIE_NAME    = "session"
	SESSION_COOKIE_MAX_AGE = 30 * 24 * 60 * 60 * 1000

	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"

	DEFAULT_SERVER_PORT = "3001"
)

type Config struct {
	EnvName string `mapstructure:"ENV_NAME"`
	Port    string `mapstructure:"PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AuthTokenSecret   string `mapstructure:"AUTH_TOKEN_SECRET"`
	CryptedDataSecret string `mapstructure:"CRYPTED_DATA_SECRET"`

	RootUserEmail    string `mapstructure:"ROOT_USER_EMAIL"`
	RootUserPassword string `mapstructure:"ROOT_USER_PASSWORD"`

	TelemetryID string `mapstructure:"TELEMETRY_ID"`
}

var config Config

func Init(env string) {
	var err error
	viper.SetConfigType("env")
	viper.SetConfigName(env)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	viper.Unmarshal(&config)
}

func IsLive() bool {
	return config.EnvName == ENV_PRODUCTION
}

func GetConfig() *Config {
	return &config
}

func GetServerPort() string {
	if config.Port == "" {
		return DEFAULT_SERVER_PORT
	}
	return config.Port
}

func GetRootUser() (string, string) {
	return config.RootUserEmail, config.RootUserPassword
}
