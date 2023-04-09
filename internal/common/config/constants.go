package config

const (
	PAGINATION_COUNT = 20

	app_name     = "slashbase"
	app_db_file  = "app.db"
	app_env_file = ".env"

	ENV_NAME_PRODUCTION  = "production"
	ENV_NAME_DEVELOPMENT = "development"

	BUILD_DESKTOP = "desktop"
	BUILD_SERVER  = "server"

	SESSION_COOKIE_NAME    = "session"
	SESSION_COOKIE_MAX_AGE = 30 * 24 * 60 * 60 * 1000

	DEFAULT_SERVER_PORT = "3000"
)
