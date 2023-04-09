package db

import (
	"fmt"
	"os"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func InitGormDB() {
	dbPath := config.GetAppDatabaseFilePath()
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
