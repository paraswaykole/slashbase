package main

import (
	"github.com/slashbaseide/slashbase/cmd"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/server"
)

var build = config.BUILD_DEVELOPMENT
var version = config.BUILD_DEVELOPMENT
var isCli bool

func main() {
	config.Init(build, version)
	db.InitGormDB()

	isCli = config.GetConfig().Cli

	server.Init(isCli)

	if isCli {
		cmd.Execute()
	}
}
