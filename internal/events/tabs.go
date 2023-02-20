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
	eventUpdateTab             = "event:update:tab"
	eventCloseTab              = "event:close:tab"
)

func (TabsEventListeners) CreateNewTab(ctx context.Context) {
	runtime.EventsOn(ctx, eventCreateTab, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		tabType := args[2].(string)
		modelschema := args[3].(string)
		modelname := args[4].(string)
		queryID := args[5].(string)
		tab, err := tabController.CreateTab(dbConnectionId, tabType, modelschema, modelname, queryID)
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

func (TabsEventListeners) UpdateTab(ctx context.Context) {
	runtime.EventsOn(ctx, eventUpdateTab, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		tabID := args[2].(string)
		tabType := args[3].(string)
		metadata := args[4].(map[string]interface{})
		tab, err := tabController.UpdateTab(dbConnectionId, tabID, tabType, metadata)
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

func (TabsEventListeners) CloseTab(ctx context.Context) {
	runtime.EventsOn(ctx, eventCloseTab, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnID := args[1].(string)
		tabID := args[2].(string)
		err := tabController.CloseTab(dbConnID, tabID)
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
