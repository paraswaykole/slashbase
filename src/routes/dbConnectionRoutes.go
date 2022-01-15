package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/controllers"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/utils"
	"slashbase.com/backend/src/views"
)

type DBConnectionRoutes struct{}

var dbConnController controllers.DBConnectionController

func (dbcr DBConnectionRoutes) CreateDBConnection(c *gin.Context) {
	var createBody struct {
		ProjectID   string `json:"projectId"`
		Name        string `json:"name"`
		Host        string `json:"host"`
		Port        string `json:"port"`
		Password    string `json:"password"`
		User        string `json:"user"`
		DBName      string `json:"dbname"`
		LoginType   string `json:"loginType"`
		UseSSH      string `json:"useSSH"`
		SSHHost     string `json:"sshHost"`
		SSHUser     string `json:"sshUser"`
		SSHPassword string `json:"sshPassword"`
		SSHKeyFile  string `json:"sshKeyFile"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)

	if isAllowed, err := controllers.GetAuthUserHasRolesForProject(authUser, createBody.ProjectID, []string{models.ROLE_ADMIN}); err != nil || !isAllowed {
		return
	}

	dbConn, err := dbConnController.CreateDBConnection(authUser, createBody.ProjectID, createBody.Name, createBody.Host, createBody.Port,
		createBody.User, createBody.Password, createBody.DBName, createBody.LoginType, createBody.UseSSH, createBody.SSHHost, createBody.SSHUser, createBody.SSHPassword, createBody.SSHKeyFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildDBConnection(dbConn),
	})
}

func (dbcr DBConnectionRoutes) GetDBConnections(c *gin.Context) {
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)

	dbConns, err := dbConnController.GetDBConnections(authUserProjectIds)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	dbConnViews := []views.DBConnectionView{}
	for _, dbConn := range dbConns {
		dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dbConnViews,
	})
}

func (dbcr DBConnectionRoutes) GetSingleDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	dbConn, err := dbConnController.GetSingleDBConnection(authUser, dbConnID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	// TODO: check if authUser is member of project
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildDBConnection(dbConn),
	})
}

func (dbcr DBConnectionRoutes) GetDBConnectionsByProject(c *gin.Context) {
	projectID := c.Param("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	dbConns, err := dbConnController.GetDBConnectionsByProject(projectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	dbConnViews := []views.DBConnectionView{}
	for _, dbConn := range dbConns {
		dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dbConnViews,
	})
}
