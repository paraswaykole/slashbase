package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"slashbase.com/backend/internal/config"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func InitGormDB() {
	db, err = gorm.Open(sqlite.Open(config.APP_DATABASE_FILE), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
