package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/db"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/queryengines"
	"slashbase.com/backend/src/server"
	"slashbase.com/backend/src/sshtunnel"
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
	autoMigrate()
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
		&models.DBQuery{},
		&models.DBQueryLog{},
	)
	err = db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func initUnusedRemovalThreads() {
	go sshtunnel.RemoveUnusedTunnels()
	go queryengines.RemoveUnusedConnections()
}
