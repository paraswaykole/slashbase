package server

import (
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/tasks"
	"github.com/slashbaseide/slashbase/internal/server/app"
	"github.com/slashbaseide/slashbase/internal/server/setup"
)

func Start() {
	setup.SetupServer()
	tasks.InitCron()
	analytics.InitAnalytics()
	analytics.SendTelemetryEvent()

	serverApp := app.CreateFiberApp()
	app.SetupRoutes(serverApp)
	serverApp.Listen(":" + config.DEFAULT_SERVER_PORT)
}
