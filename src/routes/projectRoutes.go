package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/controllers"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/utils"
	"slashbase.com/backend/src/views"
)

type ProjectRoutes struct{}

var projectController controllers.ProjectController

func (pr ProjectRoutes) CreateProject(c *gin.Context) {
	var createBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)

	project, projectMember, err := projectController.CreateProject(authUser, createBody.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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

func (pr ProjectRoutes) GetProjects(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	projectMembers, err := projectController.GetProjects(authUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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

func (pr ProjectRoutes) GetProjectMembers(c *gin.Context) {
	projectID := c.Param("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "not allowed",
		})
		return
	}
	projectMembers, err := projectController.GetProjectMembers(projectID)
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

func (pr ProjectRoutes) AddProjectMember(c *gin.Context) {
	projectID := c.Param("projectId")
	authUser := middlewares.GetAuthUser(c)
	var addMemberBody struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}
	c.BindJSON(&addMemberBody)

	if isAllowed, err := controllers.GetAuthUserHasRolesForProject(authUser, projectID, []string{models.ROLE_ADMIN}); err != nil || !isAllowed {
		return
	}

	newProjectMember, err := projectController.AddProjectMember(projectID, addMemberBody.Email, addMemberBody.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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

func (pr ProjectRoutes) DeleteProjectMember(c *gin.Context) {
	projectId := c.Param("projectId")
	userId := c.Param("userId")
	authUser := middlewares.GetAuthUser(c)

	if isAllowed, err := controllers.GetAuthUserHasRolesForProject(authUser, projectId, []string{models.ROLE_ADMIN}); err != nil || !isAllowed {
		return
	}

	err := projectController.DeleteProjectMember(projectId, userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (pr ProjectRoutes) DeleteProject(c *gin.Context) {
	projectId := c.Param("projectId")
	authUser := middlewares.GetAuthUser(c)

	if isAllowed, err := controllers.GetAuthUserHasRolesForProject(authUser, projectId, []string{models.ROLE_ADMIN}); err != nil || !isAllowed {
		return
	}

	err := projectController.DeleteProject(projectId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
