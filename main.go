package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/internal/queryengines"
	"slashbase.com/backend/internal/server"
	"slashbase.com/backend/internal/sshtunnel"
	"slashbase.com/backend/internal/tasks"
)

func main() {
	environment := flag.String("e", "local", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.InitGormDB()
	tasks.InitCron()
	autoMigrate()
	configureRootUser()
	queryengines.InitQueryEngines()
	initUnusedRemovalThreads()
	server.Init()
}

func autoMigrate() {
	err := db.GetDB().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		os.Exit(1)
	}
	db.GetDB().AutoMigrate(
		&models.User{},
		&models.UserSession{},
		&models.Project{},
		&models.ProjectMember{},
		&models.DBConnection{},
		&models.DBConnectionUser{},
		&models.DBQuery{},
		&models.DBQueryLog{},
	)
	err = db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func configureRootUser() {
	rootUserConfig := config.GetRootUser()
	rootUser, err := models.NewUser(rootUserConfig.Email, rootUserConfig.Password)
	if err != nil {
		os.Exit(1)
	}
	rootUser.IsRoot = true
	var userDao daos.UserDao
	_, err = userDao.GetRootUserOrCreate(*rootUser)
	if err != nil {
		os.Exit(1)
	}
}

func initUnusedRemovalThreads() {
	go sshtunnel.RemoveUnusedTunnels()
	go queryengines.RemoveUnusedConnections()
}
