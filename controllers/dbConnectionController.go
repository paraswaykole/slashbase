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

type DBConnectionController struct{}

var dbConnDao daos.DBConnectionDao

func (dbcc DBConnectionController) CreateDBConnection(c *gin.Context) {
	var createBody struct {
		TeamID      string `json:"teamId"`
		Name        string `json:"name"`
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
	dbConn, err := models.NewPostgresDBConnection(authUser.ID, createBody.TeamID, createBody.Name, createBody.Host, createBody.Port,
		createBody.User, createBody.Password, createBody.DBName, createBody.UseSSH, createBody.SSHHost, createBody.SSHUser, createBody.SSHPassword, createBody.SSHKeyFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	err = dbConnDao.CreateDBConnection(dbConn)
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

func (dbcc DBConnectionController) GetDBConnections(c *gin.Context) {
	authUserTeamIds := middlewares.GetAuthUserTeamIds(c)
	var getBody struct {
		TeamIDs []string `json:"teamIds"`
	}
	c.BindJSON(&getBody)

	for _, teamID := range getBody.TeamIDs {
		if !utils.ContainsString(*authUserTeamIds, teamID) {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   errors.New("teamid " + teamID + "not allowed"),
			})
			return
		}
	}

	dbConns, err := dbConnDao.GetDBConnectionsByTeamIds(getBody.TeamIDs)
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

func (dbcc DBConnectionController) GetDBConnectionsByTeam(c *gin.Context) {
	teamID := c.Param("teamId")
	authUserTeamIds := middlewares.GetAuthUserTeamIds(c)
	if !utils.ContainsString(*authUserTeamIds, teamID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	dbConns, err := dbConnDao.GetDBConnectionsByTeam(teamID)
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
