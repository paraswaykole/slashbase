package main

import (
	"fmt"

	"slashbase.com/backend/cmd"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/server"
	"slashbase.com/backend/internal/setup"
	"slashbase.com/backend/internal/tasks"
	"slashbase.com/backend/pkg/queryengines"
)

var Build = config.ENV_DEVELOPMENT

func main() {
	config.Init(Build)
	db.InitGormDB()
	setup.SetupApp()
	queryengines.Init()
	tasks.InitCron()
	fmt.Println("Running Slashbase IDE at http://localhost:" + config.GetServerPort())
	fmt.Println("Type 'help' for more info on cli.")
	server.Init()
	cmd.Execute()
}
