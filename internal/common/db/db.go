package db

import (
	"fmt"
	"os"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"gorm.io/driver/postgres"
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
	if config.IsDesktop() {
		initGormDBDesktop()
	} else {
		initGormDBServer()
	}
}

func initGormDBDesktop() {
	dbPath := config.GetAppDatabaseFilePath()
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initGormDBServer() {
	dbConfigData := config.GetConfig().AppDB
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Kolkata", dbConfigData.Host, dbConfigData.User, dbConfigData.Pass, dbConfigData.Name, dbConfigData.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
