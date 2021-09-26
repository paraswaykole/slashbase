package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"slashbase.com/backend/src/config"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func InitGormDB() {
	dbConfig := config.GetDatabaseConfig()
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Kolkata", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
