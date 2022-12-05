package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/controllers"
	"slashbase.com/backend/internal/middlewares"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/internal/views"
)

type DBConnectionHandlers struct{}

var dbConnController controllers.DBConnectionController

func (DBConnectionHandlers) CreateDBConnection(c *gin.Context) {
	var createBody struct {
		ProjectID   string `json:"projectId"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Scheme      string `json:"scheme"`
		Host        string `json:"host"`
		Port        string `json:"port"`
		Password    string `json:"password"`
		User        string `json:"user"`
		DBName      string `json:"dbname"`
		UseSSH      string `json:"useSSH"`
		SSHHost     string `json:"sshHost"`
		SSHUser     string `json:"sshUser"`
		SSHPassword string `json:"sshPassword"`
		SSHKeyFile  string `json:"sshKeyFile"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)

	dbConn, err := dbConnController.CreateDBConnection(authUser, createBody.ProjectID, createBody.Name, createBody.Type, createBody.Scheme, createBody.Host, createBody.Port,
		createBody.User, createBody.Password, createBody.DBName, createBody.UseSSH, createBody.SSHHost, createBody.SSHUser, createBody.SSHPassword, createBody.SSHKeyFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (DBConnectionHandlers) GetDBConnections(c *gin.Context) {
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)

	dbConns, err := dbConnController.GetDBConnections(authUserProjectIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (DBConnectionHandlers) DeleteDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	err := dbConnController.DeleteDBConnection(authUser, dbConnID)
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

func (DBConnectionHandlers) GetSingleDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	dbConn, err := dbConnController.GetSingleDBConnection(authUser, dbConnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (DBConnectionHandlers) GetDBConnectionsByProject(c *gin.Context) {
	projectID := c.Param("projectId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	if !utils.ContainsString(*authUserProjectIds, projectID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	dbConns, err := dbConnController.GetDBConnectionsByProject(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
