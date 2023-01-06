package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const (
	PAGINATION_COUNT = 20

	app_name    = "slashbase"
	app_db_file = "app.db"

	BUILD_PRODUCTION  = "production"
	BUILD_DEVELOPMENT = "development"

	DEFAULT_SERVER_PORT = "22022"
)

func GetAppDatabaseFilePath() string {
	if !IsLive() {
		return app_db_file
	}
	var filePath string
	if runtime.GOOS == "windows" {
		// Get the %LOCALAPPDATA% path
		localAppData := os.Getenv("LOCALAPPDATA")
		// Set the file name and path
		filePath = filepath.Join(localAppData, app_name, app_db_file)
	} else if runtime.GOOS == "darwin" {
		// Get the user's home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		filePath = filepath.Join(homeDir, "Library", "Application Support", app_name, app_db_file)
	} else if runtime.GOOS == "linux" {
		filePath = filepath.Join("/usr/local", app_name, app_db_file)
	} else {
		panic(errors.New("not implemented"))
	}
	err := os.MkdirAll(filepath.Dir(filePath), 0700)
	if err != nil {
		panic(err)
	}
	return filePath
}
