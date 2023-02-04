package events

import (
	"context"
	"time"

	"github.com/slashbaseide/slashbase/internal/analytics"
	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/slashbaseide/slashbase/internal/utils"
	"github.com/slashbaseide/slashbase/internal/views"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type QueryEventListeners struct{}

var queryController controllers.QueryController

const (
	eventRunQuery                      = "event:run:query"
	eventGetData                       = "event:get:data"
	eventGetDataModels                 = "event:get:datamodels"
	eventGetSingleDataModel            = "event:getsingle:datamodel"
	eventAddSingleDataModelField       = "event:addsingle:datamodelfield"
	eventDeleteSingleDataModelField    = "event:deletesingle:datamodelfield"
	eventAddData                       = "event:add:data"
	eventDeleteData                    = "event:delete:data"
	eventUpdateSingleData              = "event:updatesingle:data"
	eventAddSingleDataModelIndex       = "event:addsingle:datamodelindex"
	eventDeleteSingleDataModelIndex    = "event:deletesingle:datamodelindex"
	eventSaveDBQuery                   = "event:save:dbquery"
	eventDeleteDBQuery                 = "event:delete:dbquery"
	eventGetDBQueriesInDBConnection    = "event:get:dbqueries:indbconnection"
	eventGetSingleDBQuery              = "event:getsingle:dbquery"
	eventGetQueryHistoryInDBConnection = "event:get:queryhistory:indbconnection"
)

func (QueryEventListeners) RunQuery(ctx context.Context) {
	runtime.EventsOn(ctx, eventRunQuery, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnectionId := args[1].(string)
		query := args[2].(string)
		analytics.SendRunQueryEvent()
		response, err := queryController.RunQuery(dbConnectionId, query)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    response,
		})
	})
}

func (QueryEventListeners) GetData(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetData, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		var filter, sort []string
		if data["filter"] != nil {
			filter = utils.InterfaceArrayToStringArray(data["filter"].([]interface{}))
		}
		if data["sort"] != nil {
			sort = utils.InterfaceArrayToStringArray(data["sort"].([]interface{}))
		}
		analytics.SendLowCodeDataViewEvent()
		responsedata, err := queryController.GetData(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["fetchCount"].(bool), int(data["limit"].(float64)), int64(data["offset"].(float64)), filter, sort)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responsedata,
		})
	})
}

func (QueryEventListeners) GetDataModels(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetDataModels, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnId := args[1].(string)
		dataModels, err := queryController.GetDataModels(dbConnId)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    dataModels,
		})
	})
}

func (QueryEventListeners) GetSingleDataModel(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetSingleDataModel, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnId := args[1].(string)
		schema := args[2].(string)
		name := args[3].(string)
		analytics.SendLowCodeModelViewEvent()
		data, err := queryController.GetSingleDataModel(dbConnId, schema, name)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    data,
		})
	})
}

func (QueryEventListeners) AddSingleDataModelField(ctx context.Context) {
	runtime.EventsOn(ctx, eventAddSingleDataModelField, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		responseData, err := queryController.AddSingleDataModelField(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["fieldName"].(string), data["dataType"].(string))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) DeleteSingleDataModelField(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteSingleDataModelField, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		responseData, err := queryController.DeleteSingleDataModelField(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["fieldName"].(string))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) AddData(ctx context.Context) {
	runtime.EventsOn(ctx, eventAddData, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		responseData, err := queryController.AddData(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["data"].(map[string]interface{}))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) DeleteData(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteData, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		var ids []string
		if data["ids"] != nil {
			ids = utils.InterfaceArrayToStringArray(data["ids"].([]interface{}))
		}
		responseData, err := queryController.DeleteData(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), ids)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) UpdateSingleData(ctx context.Context) {
	runtime.EventsOn(ctx, eventUpdateSingleData, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		responseData, err := queryController.UpdateSingleData(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["id"].(string), data["columnName"].(string), data["value"].(string))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) AddSingleDataModelIndex(ctx context.Context) {
	runtime.EventsOn(ctx, eventAddSingleDataModelIndex, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		var fieldNames []string
		if data["fieldNames"] != nil {
			fieldNames = utils.InterfaceArrayToStringArray(data["fieldNames"].([]interface{}))
		}
		responseData, err := queryController.AddSingleDataModelIndex(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["indexName"].(string), fieldNames, data["isUnique"].(bool))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) DeleteSingleDataModelIndex(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteSingleDataModelIndex, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		responseData, err := queryController.DeleteSingleDataModelIndex(data["dbConnectionId"].(string), data["schema"].(string), data["name"].(string), data["indexName"].(string))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    responseData,
		})
	})
}

func (QueryEventListeners) SaveDBQuery(ctx context.Context) {
	runtime.EventsOn(ctx, eventSaveDBQuery, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		data := args[1].(map[string]interface{})
		analytics.SendSavedQueryEvent()
		queryObj, err := queryController.SaveDBQuery(data["dbConnectionId"].(string), data["name"].(string), data["query"].(string), data["queryId"].(string))
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildDBQueryView(queryObj),
		})
	})
}

func (QueryEventListeners) DeleteDBQuery(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteDBQuery, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		queryId := args[1].(string)
		err := queryController.DeleteDBQuery(queryId)
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

func (QueryEventListeners) GetDBQueriesInDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetDBQueriesInDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnID := args[1].(string)
		dbQueries, err := queryController.GetDBQueriesInDBConnection(dbConnID)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		dbQueryViews := []views.DBQueryView{}
		for _, dbQuery := range dbQueries {
			dbQueryViews = append(dbQueryViews, *views.BuildDBQueryView(dbQuery))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    dbQueryViews,
		})
	})
}

func (QueryEventListeners) GetSingleDBQuery(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetSingleDBQuery, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		queryID := args[1].(string)
		dbQuery, err := queryController.GetSingleDBQuery(queryID)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildDBQueryView(dbQuery),
		})
	})
}

func (QueryEventListeners) GetQueryHistoryInDBConnection(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetQueryHistoryInDBConnection, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		dbConnID := args[1].(string)
		beforeInt, isExist := args[2].(float64)
		var before time.Time
		if !isExist {
			before = time.Now()
		} else {
			before = utils.UnixNanoToTime(int64(beforeInt))
		}
		dbQueryLogs, next, err := queryController.GetQueryHistoryInDBConnection(dbConnID, before)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		dbQueryLogViews := []views.DBQueryLogView{}
		for _, dbQueryLog := range dbQueryLogs {
			dbQueryLogViews = append(dbQueryLogViews, *views.BuildDBQueryLogView(dbQueryLog))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"list": dbQueryLogViews,
				"next": next,
			},
		})
	})
}
