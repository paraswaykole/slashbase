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
	queryengines.Init()
	tasks.InitCron()
	server.Init()
}
