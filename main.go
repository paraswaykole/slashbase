package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
	"slashbase.com/backend/server"
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
	server.Init()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(&models.User{}, &models.UserSession{}, &models.Team{}, &models.TeamMember{})
	err := db.GetDB().SetupJoinTable(&models.User{}, "Teams", &models.TeamMember{})
	if err != nil {
		os.Exit(1)
	}
}
