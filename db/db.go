package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"slashbase.com/backend/config"
	"slashbase.com/backend/models/user"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}

func InitGormDB() {
	nConfig := config.GetConfig()
	host := nConfig.GetString("database.host")
	user := nConfig.GetString("database.user")
	password := nConfig.GetString("database.password")
	database := nConfig.GetString("database.database")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Kolkata", host, user, password, database)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	autoMigrate()
}

func autoMigrate() {
	db.AutoMigrate(&user.User{})
}
