package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/slashbaseide/slashbase/internal/common/utils"
)

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
