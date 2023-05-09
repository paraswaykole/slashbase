package main

import (
	"embed"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/desktop"
	"github.com/slashbaseide/slashbase/internal/server"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
)

//go:embed all:frontend/dist
var assets embed.FS

var build = config.BUILD_DESKTOP
var envName = config.ENV_NAME_DEVELOPMENT
var version = config.ENV_NAME_DEVELOPMENT

func main() {
	config.Init(build, envName, version)
	db.InitGormDB()
	queryengines.Init()
	if config.IsDesktop() {
		desktop.Start(assets)
	} else {
		server.Start(assets)
	}
}
