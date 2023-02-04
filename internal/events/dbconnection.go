package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/slashbaseide/slashbase/internal/views"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DBConnectionEventListeners struct{}

var dbConnController controllers.DBConnectionController

const (
	eventCreateDBConnection        = "event:create:dbconnection"
	eventGetDBConnections          = "event:get:dbconnections"
	eventDeleteDBConnection        = "event:delete:dbconnection"
	eventGetSingleDBConnection     = "event:getsingle:dbconnection"
	eventGetDBConnectionsByProject = "event:get:dbconnections:byproject"
)

func (DBConnectionEventListeners) CreateDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventCreateDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		dbConn, err := dbConnController.CreateDBConnection(data["projectId"].(string), data["name"].(string), data["type"].(string), data["scheme"].(string), data["host"].(string), data["port"].(string),
			data["user"].(string), data["password"].(string), data["dbname"].(string), data["useSSH"].(string), data["sshHost"].(string), data["sshUser"].(string), data["sshPassword"].(string), data["sshKeyFile"].(string), data["useSSL"].(bool))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildDBConnection(dbConn),
		})
	})
}

func (DBConnectionEventListeners) GetDBConnections(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetDBConnections, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConns, err := dbConnController.GetDBConnections()
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		dbConnViews := []views.DBConnectionView{}
		for _, dbConn := range dbConns {
			dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    dbConnViews,
		})
	})
}

func (DBConnectionEventListeners) DeleteDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnID := args[1].(string)
		err := dbConnController.DeleteDBConnection(dbConnID)
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

func (DBConnectionEventListeners) GetSingleDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetSingleDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnID := args[1].(string)
		dbConn, err := dbConnController.GetSingleDBConnection(dbConnID)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildDBConnection(dbConn),
		})
	})
}

func (DBConnectionEventListeners) GetDBConnectionsByProject(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetDBConnectionsByProject, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		projectID := args[1].(string)
		dbConns, err := dbConnController.GetDBConnectionsByProject(projectID)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		dbConnViews := []views.DBConnectionView{}
		for _, dbConn := range dbConns {
			dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    dbConnViews,
		})
	})
}
