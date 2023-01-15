package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/slashbaseide/slashbase/internal/utils"
)

var config AppConfig

func Init(buildName, version string) {
	if buildName == BUILD_DEVELOPMENT {
		err := godotenv.Load("development.env")
		if err != nil {
			log.Fatal("Error loading development.env file")
		}
	} else if buildName == BUILD_PRODUCTION {
		err := godotenv.Load(GetAppEnvFilePath())
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	config = newConfig(buildName, version)
}

func IsLive() bool {
	return config.BuildName == BUILD_PRODUCTION
}

func GetConfig() *AppConfig {
	return &config
}

func GetAppEnvFilePath() string {
	var filePath string
	if runtime.GOOS == "windows" {
		// Get the %LOCALAPPDATA% path
		localAppData := os.Getenv("LOCALAPPDATA")
		// Set the file name and path
		filePath = filepath.Join(localAppData, app_name, app_env_file)
	} else if runtime.GOOS == "darwin" {
		// Get the user's home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		filePath = filepath.Join(homeDir, "Library", "Application Support", app_name, app_env_file)
	} else if runtime.GOOS == "linux" {
		filePath = filepath.Join("/usr/local", app_name, app_env_file)
	} else {
		panic(errors.New("not implemented"))
	}
	err := os.MkdirAll(filepath.Dir(filePath), 0700)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = createEnvFile(filePath)
		if err != nil {
			panic(err)
		}
	}
	return filePath
}

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

func createEnvFile(filePath string) error {
	hex, err := utils.RandomHex(32)
	if err != nil {
		return err
	}
	envFileData := fmt.Sprintf(`CRYPTED_DATA_SECRET=%s`, hex)
	err = os.WriteFile(filePath, []byte(envFileData), 0700)
	if err != nil {
		return err
	}
	return nil
}
