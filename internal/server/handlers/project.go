package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
	"github.com/slashbaseide/slashbase/internal/common/views"
)

type ProjectHandlers struct{}

var projectController controllers.ProjectController

func (ProjectHandlers) CreateProject(c *fiber.Ctx) error {
	var createBody struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	project, err := projectController.CreateProject(createBody.Name)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildProject(project),
	})
}

func (ProjectHandlers) GetProjects(c *fiber.Ctx) error {
	projects, err := projectController.GetProjects()
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	projectViews := []views.ProjectView{}
	for _, p := range *projects {
		projectViews = append(projectViews, views.BuildProject(&p))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    projectViews,
	})
}

func (ProjectHandlers) DeleteProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	err := projectController.DeleteProject(projectID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}
