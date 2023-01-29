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
}
