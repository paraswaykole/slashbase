package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/controllers"
	"slashbase.com/backend/internal/middlewares"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/internal/views"
)

type ProjectHandlers struct{}

var projectController controllers.ProjectController

func (ProjectHandlers) CreateProject(c *gin.Context) {
	var createBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)

	project, projectMember, err := projectController.CreateProject(authUser, createBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildProject(project, projectMember),
	})
}

func (ProjectHandlers) GetProjects(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	projectMembers, err := projectController.GetProjects(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	projectViews := []views.ProjectView{}
	for _, t := range *projectMembers {
		projectViews = append(projectViews, views.BuildProject(&t.Project, &t))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    projectViews,
	})
}

func (ProjectHandlers) GetProjectMembers(c *gin.Context) {
	projectID := c.Param("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "not allowed",
		})
		return
	}
	projectMembers, err := projectController.GetProjectMembers(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	projectMemberViews := []views.ProjectMemberView{}
	for _, t := range *projectMembers {
		projectMemberViews = append(projectMemberViews, views.BuildProjectMember(&t))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    projectMemberViews,
	})
}

func (ProjectHandlers) AddProjectMember(c *gin.Context) {
	projectID := c.Param("projectId")
	authUser := middlewares.GetAuthUser(c)
	var addMemberBody struct {
		Email  string `json:"email"`
		RoleID string `json:"roleId"`
	}
	c.BindJSON(&addMemberBody)

	newProjectMember, err := projectController.AddProjectMember(authUser, projectID, addMemberBody.Email, addMemberBody.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildProjectMember(newProjectMember),
	})
}

func (ProjectHandlers) DeleteProjectMember(c *gin.Context) {
	projectId := c.Param("projectId")
	userId := c.Param("userId")
	authUser := middlewares.GetAuthUser(c)

	err := projectController.DeleteProjectMember(authUser, projectId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ProjectHandlers) DeleteProject(c *gin.Context) {
	projectId := c.Param("projectId")
	authUser := middlewares.GetAuthUser(c)

	err := projectController.DeleteProject(authUser, projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
