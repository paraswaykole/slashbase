package config

const (
	VERSION = "v1.0.0-beta"

	PAGINATION_COUNT = 20

	SESSION_COOKIE_NAME    = "session"
	SESSION_COOKIE_MAX_AGE = 30 * 24 * 60 * 60 * 1000

	APP_DATABASE_FILE = "data/app.db"

	ENV_PRODUCTION  = "production"
	ENV_DEVELOPMENT = "development"

	DEFAULT_SERVER_PORT = "3000"
)
