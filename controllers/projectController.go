package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type ProjectController struct{}

var projectDao daos.ProjectDao

func (tc ProjectController) CreateProject(c *gin.Context) {
	var createBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)
	project := models.NewProject(authUser, createBody.Name)
	err := projectDao.CreateProject(project)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildProject(project),
	})
}

func (tc ProjectController) GetProjects(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	projectMembers, err := projectDao.GetProjectMembersForUser(authUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	projectViews := []views.ProjectView{}
	for _, t := range *projectMembers {
		projectViews = append(projectViews, views.BuildProject(&t.Project))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    projectViews,
	})
}

func (tc ProjectController) GetProjectMembers(c *gin.Context) {
	projectID := c.Param("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}
	projectMembers, err := projectDao.GetProjectMembers(projectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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
