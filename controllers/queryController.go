package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/queryengines"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type QueryController struct{}

func (qc QueryController) RunQuery(c *gin.Context) {
	var runBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	c.BindJSON(&runBody)

	dbConn, err := dbConnDao.GetDBConnectionByID(runBody.DBConnectionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}

	if isAllowed, err := middlewares.GetAuthUserHasRolesForProject(c, dbConn.ProjectID, []string{models.ROLE_ADMIN, models.ROLE_DEVELOPER}); err != nil || !isAllowed {
		return
	}

	data, err := queryengines.RunQuery(dbConn, runBody.Query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (qc QueryController) GetData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	schema := c.Query("schema")
	name := c.Query("name")
	fetchCount := c.Query("count") == "true"
	limit := 200
	offsetStr := c.Query("offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = int64(0)
	}

	authUserProjects := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	if !utils.ContainsString(*authUserProjects, dbConn.ProjectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Not allowed to run query",
		})
		return
	}

	data, err := queryengines.GetData(dbConn, schema, name, limit, offset, fetchCount)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (qc QueryController) GetDataModels(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	authUserProjects := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	if !utils.ContainsString(*authUserProjects, dbConn.ProjectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Not allowed to run query",
		})
		return
	}

	dataModels, err := queryengines.GetDataModels(dbConn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	data := []*views.DBDataModel{}
	for _, table := range dataModels["rows"].([]map[string]interface{}) {
		view := views.BuildDBDataModel(dbConn, table)
		if view != nil {
			data = append(data, view)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (qc QueryController) UpdateSingleData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	var updateBody struct {
		Schema     string `json:"schema"`
		Name       string `json:"name"`
		CTID       string `json:"ctid"`
		ColumnName string `json:"columnName"`
		Value      string `json:"value"`
	}
	c.BindJSON(&updateBody)

	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}

	if isAllowed, err := middlewares.GetAuthUserHasRolesForProject(c, dbConn.ProjectID, []string{models.ROLE_ADMIN, models.ROLE_DEVELOPER}); err != nil || !isAllowed {
		return
	}

	data, err := queryengines.UpdateSingleData(dbConn, updateBody.Schema, updateBody.Name, updateBody.CTID, updateBody.ColumnName, updateBody.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
