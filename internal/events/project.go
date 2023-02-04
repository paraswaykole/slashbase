package events

import (
	"context"

	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/slashbaseide/slashbase/internal/views"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ProjectEventListeners struct{}

var projectController controllers.ProjectController

const (
	eventCreateProject = "event:create:project"
	eventGetProjects   = "event:get:projects"
	eventDeleteProject = "event:delete:project"
)

func (ProjectEventListeners) CreateProject(ctx context.Context) {
	runtime.EventsOn(ctx, eventCreateProject, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		projectName := args[1].(string)
		project, err := projectController.CreateProject(projectName)
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    views.BuildProject(project),
		})
	})
}

func (ProjectEventListeners) GetProjects(ctx context.Context) {
	runtime.EventsOn(ctx, eventGetProjects, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		projects, err := projectController.GetProjects()
		if err != nil {
			runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		projectViews := []views.ProjectView{}
		for _, p := range *projects {
			projectViews = append(projectViews, views.BuildProject(&p))
		}
		runtime.EventsEmit(ctx, responseEventName, map[string]interface{}{
			"success": true,
			"data":    projectViews,
		})
	})
}

func (ProjectEventListeners) DeleteProject(ctx context.Context) {
	runtime.EventsOn(ctx, eventDeleteProject, func(args ...interface{}) {
		responseEventName := args[0].(string)
		defer recovery(ctx, responseEventName)
		projectID := args[1].(string)
		err := projectController.DeleteProject(projectID)
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
