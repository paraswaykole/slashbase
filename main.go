package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
	"slashbase.com/backend/queryengines"
	"slashbase.com/backend/server"
	"slashbase.com/backend/sshtunnel"
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
	initUnusedRemovalThreads()
	server.Init()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(&models.User{}, &models.UserSession{}, &models.Project{}, &models.ProjectMember{}, &models.DBConnection{})
	err := db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func initUnusedRemovalThreads() {
	go sshtunnel.RemoveUnusedTunnels()
	go queryengines.RemoveUnusedConnections()
}
