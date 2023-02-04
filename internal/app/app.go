package app

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/analytics"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
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
	setupEvents(ctx)
	analytics.SendTelemetryEvent()
}

// AppID returns unqiue appid.
func (a *App) AppID() string {
	setting, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_APP_ID)
	if err != nil {
		return ""
	}
	return setting.Value
}
