package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/slashbaseide/slashbase/internal/views"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type TabsEventListeners struct{}

var tabController controllers.TabsController

const (
	eventCreateTab             = "event:create:tab"
	eventGetTabsByDBConnection = "event:get:tabs:bydbconnection"
)

func (TabsEventListeners) CreateNewTab(ctx context.Context) {
	runtime.EventsOn(ctx, eventCreateTab, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		// TODO: handle more args...
		tab, err := tabController.CreateTab(dbConnectionId)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildTabView(tab),
		})
	})
}

func (TabsEventListeners) GetTabsByDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetTabsByDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		tabs, err := tabController.GetTabsByDBConnection(dbConnectionId)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		tabViews := []views.TabView{}
		for _, t := range *tabs {
			tabViews = append(tabViews, *views.BuildTabView(&t))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    tabViews,
		})
	})
}
