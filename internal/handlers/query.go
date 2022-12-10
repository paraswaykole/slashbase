package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/controllers"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/internal/views"
)

type QueryHandlers struct{}

var queryController controllers.QueryController

func (QueryHandlers) RunQuery(c *gin.Context) {
	var runBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	c.BindJSON(&runBody)

	data, err := queryController.RunQuery(runBody.DBConnectionID, runBody.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) GetData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	schema := c.Query("schema")
	name := c.Query("name")
	fetchCount := c.Query("count") == "true"
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}
	offsetStr := c.Query("offset")
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = int64(0)
	}
	filter, _ := c.GetQueryArray("filter[]")
	sort, _ := c.GetQueryArray("sort[]")

	data, err := queryController.GetData(dbConnId, schema, name, fetchCount, limit, offset, filter, sort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) GetDataModels(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	dataModels, err := queryController.GetDataModels(dbConnId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dataModels,
	})
}

func (QueryHandlers) GetSingleDataModel(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	schema := c.Query("schema")
	name := c.Query("name")

	data, err := queryController.GetSingleDataModel(dbConnId, schema, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) AddSingleDataModelField(c *gin.Context) {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
		DataType       string `json:"dataType"`
	}
	c.BindJSON(&reqBody)

	data, err := queryController.AddSingleDataModelField(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName, reqBody.DataType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) DeleteSingleDataModelField(c *gin.Context) {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
	}
	c.BindJSON(&reqBody)

	data, err := queryController.DeleteSingleDataModelField(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) AddData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	var addBody struct {
		Schema string                 `json:"schema"`
		Name   string                 `json:"name"`
		Data   map[string]interface{} `json:"data"`
	}
	c.BindJSON(&addBody)

	data, err := queryController.AddData(dbConnId, addBody.Schema, addBody.Name, addBody.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) DeleteData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	var deleteBody struct {
		Schema string   `json:"schema"`
		Name   string   `json:"name"`
		IDs    []string `json:"ids"` // ctid for postgres, _id for mongo
	}
	c.BindJSON(&deleteBody)

	data, err := queryController.DeleteData(dbConnId, deleteBody.Schema, deleteBody.Name, deleteBody.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) UpdateSingleData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	var updateBody struct {
		Schema     string `json:"schema"`
		Name       string `json:"name"`
		ID         string `json:"id"`
		ColumnName string `json:"columnName"`
		Value      string `json:"value"`
	}
	c.BindJSON(&updateBody)

	data, err := queryController.UpdateSingleData(dbConnId, updateBody.Schema, updateBody.Name, updateBody.ID, updateBody.ColumnName, updateBody.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) AddSingleDataModelIndex(c *gin.Context) {
	var reqBody struct {
		DBConnectionID string   `json:"dbConnectionId"`
		Schema         string   `json:"schema"`
		Name           string   `json:"name"`
		IndexName      string   `json:"indexName"`
		FieldNames     []string `json:"fieldNames"`
		IsUnique       bool     `json:"isUnique"`
	}
	c.BindJSON(&reqBody)

	data, err := queryController.AddSingleDataModelIndex(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName, reqBody.FieldNames, reqBody.IsUnique)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) DeleteSingleDataModelIndex(c *gin.Context) {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		IndexName      string `json:"indexName"`
	}
	c.BindJSON(&reqBody)

	data, err := queryController.DeleteSingleDataModelIndex(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) SaveDBQuery(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	var createBody struct {
		Name    string `json:"name"`
		Query   string `json:"query"`
		QueryID string `json:"queryId"`
	}
	c.BindJSON(&createBody)

	queryObj, err := queryController.SaveDBQuery(dbConnId, createBody.Name, createBody.Query, createBody.QueryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildDBQueryView(queryObj),
	})
}

func (QueryHandlers) DeleteDBQuery(c *gin.Context) {
	queryID := c.Param("queryId")

	err := queryController.DeleteDBQuery(queryID)
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

func (QueryHandlers) GetDBQueriesInDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")

	dbQueries, err := queryController.GetDBQueriesInDBConnection(dbConnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	dbQueryViews := []views.DBQueryView{}
	for _, dbQuery := range dbQueries {
		dbQueryViews = append(dbQueryViews, *views.BuildDBQueryView(dbQuery))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dbQueryViews,
	})
}

func (QueryHandlers) GetSingleDBQuery(c *gin.Context) {
	queryID := c.Param("queryId")

	dbQuery, err := queryController.GetSingleDBQuery(queryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildDBQueryView(dbQuery),
	})
}

func (QueryHandlers) GetQueryHistoryInDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")

	beforeInt, err := strconv.ParseInt(c.Query("before"), 10, 64)
	var before time.Time
	if err != nil {
		before = time.Now()
	} else {
		before = utils.UnixNanoToTime(beforeInt)
	}

	dbQueryLogs, next, err := queryController.GetQueryHistoryInDBConnection(dbConnID, before)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	dbQueryLogViews := []views.DBQueryLogView{}
	for _, dbQueryLog := range dbQueryLogs {
		dbQueryLogViews = append(dbQueryLogViews, *views.BuildDBQueryLogView(dbQueryLog))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"list": dbQueryLogViews,
			"next": next,
		},
	})
}
