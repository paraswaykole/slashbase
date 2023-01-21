package app

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/config"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Version returns a greeting for the given name
func (a *App) Version() string {
	return config.GetConfig().Version
}
