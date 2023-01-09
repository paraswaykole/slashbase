package main

import (
	"github.com/slashbaseide/slashbase/cmd"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/db"
)

var build = config.BUILD_DEVELOPMENT
var version = config.BUILD_DEVELOPMENT

func main() {
	config.Init(build, version)
	db.InitGormDB()
	cmd.Execute()
}
