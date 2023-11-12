package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
	"github.com/slashbaseide/slashbase/internal/server/views"
)

type ProjectHandlers struct{}

var projectController controllers.ProjectController

func (ProjectHandlers) CreateProject(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var createBody struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return fiber.ErrBadRequest
	}
	project, projectMember, err := projectController.CreateProject(authUser, createBody.Name)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildProject(project, projectMember),
	})
}

func (ProjectHandlers) GetProjects(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	projectMembers, err := projectController.GetProjects(authUser)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	projectViews := []views.ProjectView{}
	for _, pm := range *projectMembers {
		projectViews = append(projectViews, views.BuildProject(&pm.Project, &pm))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    projectViews,
	})
}

func (ProjectHandlers) DeleteProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	authUser := middlewares.GetAuthUser(c)
	err := projectController.DeleteProject(authUser, projectID)
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

func (ProjectHandlers) GetProjectMembers(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   "not allowed",
		})
	}
	projectMembers, err := projectController.GetProjectMembers(projectID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	projectMemberViews := []views.ProjectMemberView{}
	for _, t := range *projectMembers {
		projectMemberViews = append(projectMemberViews, views.BuildProjectMember(&t))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    projectMemberViews,
	})
}

func (ProjectHandlers) AddProjectMember(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	authUser := middlewares.GetAuthUser(c)
	var addMemberBody struct {
		Email  string `json:"email"`
		RoleID string `json:"roleId"`
	}
	if err := c.BodyParser(&addMemberBody); err != nil {
		return fiber.ErrBadRequest
	}
	newProjectMember, err := projectController.AddProjectMember(authUser, projectID, addMemberBody.Email, addMemberBody.RoleID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildProjectMember(newProjectMember),
	})
}

func (ProjectHandlers) DeleteProjectMember(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	userId := c.Params("userId")
	authUser := middlewares.GetAuthUser(c)

	err := projectController.DeleteProjectMember(authUser, projectId, userId)
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
