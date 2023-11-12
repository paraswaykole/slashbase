package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type AIEventListeners struct{}

var aiController controllers.AIController

const (
	eventAIGenSQL     = "event:ai:gensql"
	eventAIListModels = "event:ai:listmodels"
)

func (AIEventListeners) GenSQLEvent(ctx context.Context) {
	runtime.EventsOn(ctx, eventAIGenSQL, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		text := args[2].(string)
		analytics.SendAISQLGeneratedEvent()
		output, err := aiController.GenerateSQL(dbConnectionId, text)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    output,
		})
	})
}

func (AIEventListeners) ListSupportedAIModelsEvent(ctx context.Context) {
	runtime.EventsOn(ctx, eventAIListModels, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		output := aiController.GetModels()
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    output,
		})
	})
}
