package main

import (
	"fmt"

	"github.com/slashbaseide/slashbase/cmd"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/server"
	"github.com/slashbaseide/slashbase/internal/setup"
	"github.com/slashbaseide/slashbase/internal/tasks"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
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
