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
	var createCmd struct {
		TeamID   string `json:"teamId"`
		Name     string `json:"name"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		User     string `json:"user"`
		DBName   string `json:"dbname"`
	}
	c.BindJSON(&createCmd)
	authUser := middlewares.GetAuthUser(c)
	dbConn := models.NewPostgresDBConnection(authUser.ID, createCmd.TeamID, createCmd.Name, createCmd.Host, createCmd.Port, createCmd.User, createCmd.Password, createCmd.DBName)
	err := dbConnDao.CreateDBConnection(dbConn)
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
	return
}

func (dbcc DBConnectionController) GetDBConnectionsByTeam(c *gin.Context) {
	teamID := c.Param("teamId")
	authUserTeams := middlewares.GetAuthUserTeamIds(c)
	if !utils.ContainsString(*authUserTeams, teamID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("Not allowed"),
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
	return
}
