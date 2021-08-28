package controllers

import (
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
	projectMember, err := projectDao.CreateProject(project)
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
		projectViews = append(projectViews, views.BuildProject(&t.Project, &t))
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
			"error":   "not allowed",
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

func (tc ProjectController) AddProjectMembers(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	projectID := c.Param("projectId")
	var addMemberBody struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}
	c.BindJSON(&addMemberBody)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "not allowed",
		})
		return
	}
	authUserProjectMember, err := projectDao.GetUserProjectMembersForProject(projectID, authUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}

	if authUserProjectMember.Role != models.ROLE_ADMIN {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "not allowed",
		})
		return
	}

	toAddUser, err := userDao.GetUserByEmail(addMemberBody.Email)
	if err != nil {
		// TODO: Create user and send email if doesn't exist in users table.
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "user does not exist",
		})
		return
	}

	newProjectMember, err := models.NewProjectMember(toAddUser.ID, projectID, addMemberBody.Role)
	if err != nil {
		// TODO: Create user and send email if doesn't exist in users table.
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	err = projectDao.CreateProjectMember(newProjectMember)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	newProjectMember.User = *toAddUser
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildProjectMember(newProjectMember),
	})
}
