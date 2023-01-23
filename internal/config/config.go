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

func getAppDataPath() string {
	var filePath string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		filePath = filepath.Join(homeDir, "AppData", "Local", app_name)
	} else if runtime.GOOS == "darwin" {
		filePath = filepath.Join(homeDir, "Library", "Application Support", app_name)
	} else if runtime.GOOS == "linux" {
		filePath = filepath.Join(homeDir, "."+app_name)
	} else {
		panic(errors.New("not implemented"))
	}
	return filePath
}

func GetAppEnvFilePath() string {
	filePath := filepath.Join(getAppDataPath(), app_env_file)
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
	filePath := filepath.Join(getAppDataPath(), app_db_file)
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
