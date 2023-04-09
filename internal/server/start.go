package server

import (
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/server/app"
)

func Start() {
	app := app.CreateFiberApp()

	app.Listen(":" + config.DEFAULT_SERVER_PORT)
}
