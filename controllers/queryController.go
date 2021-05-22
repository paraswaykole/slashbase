package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/queryengines"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type QueryController struct{}

func (qc QueryController) RunQuery(c *gin.Context) {
	var runCmd struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	c.BindJSON(&runCmd)
	authUserTeams := middlewares.GetAuthUserTeamIds(c)

	dbConn, err := dbConnDao.GetDBConnectionByID(runCmd.DBConnectionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if !utils.ContainsString(*authUserTeams, dbConn.TeamID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("Not allowed to run query"),
		})
		return
	}

	data, err := queryengines.RunQuery(dbConn, runCmd.Query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (qc QueryController) GetTables(c *gin.Context) {
	var runCmd struct {
		DBConnectionID string `json:"dbConnectionId"`
	}
	c.BindJSON(&runCmd)
	authUserTeams := middlewares.GetAuthUserTeamIds(c)

	dbConn, err := dbConnDao.GetDBConnectionByID(runCmd.DBConnectionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	if !utils.ContainsString(*authUserTeams, dbConn.TeamID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("Not allowed to run query"),
		})
		return
	}

	tablesData, err := queryengines.GetTables(dbConn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	data := []*views.DBTableView{}
	for _, table := range tablesData["rows"].([]interface{}) {
		view := views.BuildDBTableView(dbConn, table.(map[string]interface{}))
		if view != nil {
			data = append(data, view)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
	return
}
