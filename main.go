package main

import (
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
	server.Init()
	cmd.Execute()
}
