package server

import (
	"embed"

	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/tasks"
	"github.com/slashbaseide/slashbase/internal/server/app"
	"github.com/slashbaseide/slashbase/internal/server/setup"
)

func Start(assets embed.FS) {
	setup.SetupServer()
	analytics.InitAnalytics()
	tasks.InitCron()

	serverApp := app.CreateFiberApp()
	app.SetupRoutes(serverApp, assets)
	serverApp.Listen(":" + config.DEFAULT_SERVER_PORT)
}
