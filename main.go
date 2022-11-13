package main

import (
	"flag"
	"fmt"
	"os"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/server"
	"slashbase.com/backend/internal/setup"
	"slashbase.com/backend/internal/tasks"
	"slashbase.com/backend/pkg/queryengines"
	"slashbase.com/backend/pkg/sshtunnel"
)

func main() {
	environment := flag.String("e", config.ENV_DEVELOPMENT, "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.InitGormDB()
	setup.SetupApp()
	tasks.InitCron()
	// TODO: to be moved to cron
	queryengines.Init()
	initUnusedRemovalThreads()
	server.Init()
}

// TODO: to be moved to cron
func initUnusedRemovalThreads() {
	go sshtunnel.RemoveUnusedTunnels()
	go queryengines.RemoveUnusedConnections()
}
