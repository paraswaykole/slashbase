package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slashbaseide/slashbase/internal/controllers"
	"github.com/slashbaseide/slashbase/internal/views"
)

type ProjectHandlers struct{}

var projectController controllers.ProjectController

func (ProjectHandlers) CreateProject(c *gin.Context) {
	var createBody struct {
		Name string `json:"name"`
	}
	err := c.BindJSON(&createBody)
	if err != nil {
		return
	}

	project, err := projectController.CreateProject(createBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (ProjectHandlers) GetProjects(c *gin.Context) {

	projects, err := projectController.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	var projectViews []views.ProjectView
	for _, p := range *projects {
		projectViews = append(projectViews, views.BuildProject(&p))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    projectViews,
	})
}

func (ProjectHandlers) DeleteProject(c *gin.Context) {
	projectId := c.Param("projectId")

	err := projectController.DeleteProject(projectId)
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
