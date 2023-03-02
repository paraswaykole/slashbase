package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ConsoleEventListeners struct{}

var consoleController controllers.ConsoleController

const (
	eventRunCommand = "event:run:cmd"
)

func (ConsoleEventListeners) RunCommandEvent(ctx context.Context) {
	runtime.EventsOn(ctx, eventRunCommand, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		cmdString := args[2].(string)
		output := consoleController.RunCommand(dbConnectionId, cmdString)
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    output,
		})
	})
}
