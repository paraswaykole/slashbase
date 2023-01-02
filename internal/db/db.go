package db

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/slashbaseide/slashbase/internal/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func InitGormDB() {
	db, err = gorm.Open(sqlite.Open(config.APP_DATABASE_FILE), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
