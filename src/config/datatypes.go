package config

type DatabaseConfig struct {
	Host     string
	Database string
	User     string
	Password string
}

type RootUserConfig struct {
	Email    string
	Password string
}
