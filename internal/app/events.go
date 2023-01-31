package app

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/events"
)

func setupEvents(ctx context.Context) {
	if projectEventListeners := new(events.ProjectEventListeners); true {
		projectEventListeners.CreateProject(ctx)
		projectEventListeners.GetProjects(ctx)
		projectEventListeners.DeleteProject(ctx)
	}
	if dbConnectionEventListeners := new(events.DBConnectionEventListeners); true {
		dbConnectionEventListeners.CreateDBConnection(ctx)
		dbConnectionEventListeners.GetDBConnections(ctx)
		dbConnectionEventListeners.DeleteDBConnection(ctx)
		dbConnectionEventListeners.GetSingleDBConnection(ctx)
		dbConnectionEventListeners.GetDBConnectionsByProject(ctx)
	}
	if settingEventListeners := new(events.SettingEventListeners); true {
		settingEventListeners.GetSingleSetting(ctx)
		settingEventListeners.UpdateSingleSetting(ctx)
	}
	if queryEventListeners := new(events.QueryEventListeners); true {
		queryEventListeners.RunQuery(ctx)
		queryEventListeners.SaveDBQuery(ctx)
		queryEventListeners.GetDBQueriesInDBConnection(ctx)
		queryEventListeners.GetSingleDBQuery(ctx)
		queryEventListeners.DeleteDBQuery(ctx)
		queryEventListeners.GetQueryHistoryInDBConnection(ctx)
		queryEventListeners.GetData(ctx)
		queryEventListeners.UpdateSingleData(ctx)
		queryEventListeners.AddData(ctx)
		queryEventListeners.DeleteData(ctx)
		queryEventListeners.GetDataModels(ctx)
		queryEventListeners.GetSingleDataModel(ctx)
		queryEventListeners.AddSingleDataModelField(ctx)
		queryEventListeners.DeleteSingleDataModelField(ctx)
		queryEventListeners.AddSingleDataModelIndex(ctx)
		queryEventListeners.DeleteSingleDataModelIndex(ctx)
	}
}
