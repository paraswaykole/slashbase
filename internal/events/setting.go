package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SettingEventListeners struct{}

var settingController controllers.SettingController

const (
	eventGetSingleSetting    = "event:getsingle:setting"
	eventUpdateSingleSetting = "event:updatesingle:setting"
)

func (SettingEventListeners) GetSingleSetting(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetSingleSetting, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		name := args[1].(string)
		value, err := settingController.GetSingleSetting(name)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    value,
		})
	})
}

func (SettingEventListeners) UpdateSingleSetting(ctx context.Context) {
	runtime.EventsOn(ctx, eventUpdateSingleSetting, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		name := args[1].(string)
		value := args[2].(string)
		err := settingController.UpdateSingleSetting(name, value)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
		})
	})
}
