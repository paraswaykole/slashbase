package events

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func recovery(ctx context.Context, responseEventName string) {
	if r := recover(); r != nil {
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": false,
			"error":   "there was some problem",
		})
		return
	}
}
