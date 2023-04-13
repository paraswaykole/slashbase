package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/internal/common/views"
)

type QueryHandlers struct{}

var queryController controllers.QueryController

func (QueryHandlers) RunQuery(c *fiber.Ctx) error {
	var runBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	if err := c.BodyParser(&runBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	analytics.SendRunQueryEvent()
	response, err := queryController.RunQuery(runBody.DBConnectionID, runBody.Query)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    response,
	})
}

func (QueryHandlers) GetData(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")

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
	// TODO: update and fix
	// filter, _ := c.GetQueryArray("filter[]")
	// sort, _ := c.GetQueryArray("sort[]")
	analytics.SendLowCodeDataViewEvent()
	responsedata, err := queryController.GetData(dbConnId, schema, name, fetchCount, limit, offset, []string{}, []string{})
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responsedata,
	})
}

func (QueryHandlers) GetDataModels(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	dataModels, err := queryController.GetDataModels(dbConnId)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    dataModels,
	})
}

func (QueryHandlers) GetSingleDataModel(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	schema := c.Query("schema")
	name := c.Query("name")
	analytics.SendLowCodeModelViewEvent()
	data, err := queryController.GetSingleDataModel(dbConnId, schema, name)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

func (QueryHandlers) AddSingleDataModelField(c *fiber.Ctx) error {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
		DataType       string `json:"dataType"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.AddSingleDataModelField(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName, reqBody.DataType)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) DeleteSingleDataModelField(c *fiber.Ctx) error {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.DeleteSingleDataModelField(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) AddData(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	var addBody struct {
		Schema string                 `json:"schema"`
		Name   string                 `json:"name"`
		Data   map[string]interface{} `json:"data"`
	}
	if err := c.BodyParser(&addBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.AddData(dbConnId, addBody.Schema, addBody.Name, addBody.Data)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) DeleteData(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	var deleteBody struct {
		Schema string   `json:"schema"`
		Name   string   `json:"name"`
		IDs    []string `json:"ids"` // ctid for postgres, _id for mongo
	}
	if err := c.BodyParser(&deleteBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.DeleteData(dbConnId, deleteBody.Schema, deleteBody.Name, deleteBody.IDs)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) UpdateSingleData(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	var updateBody struct {
		Schema     string `json:"schema"`
		Name       string `json:"name"`
		ID         string `json:"id"`
		ColumnName string `json:"columnName"`
		Value      string `json:"value"`
	}
	if err := c.BodyParser(&updateBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.UpdateSingleData(dbConnId, updateBody.Schema, updateBody.Name, updateBody.ID, updateBody.ColumnName, updateBody.Value)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) AddSingleDataModelIndex(c *fiber.Ctx) error {
	var reqBody struct {
		DBConnectionID string   `json:"dbConnectionId"`
		Schema         string   `json:"schema"`
		Name           string   `json:"name"`
		IndexName      string   `json:"indexName"`
		FieldNames     []string `json:"fieldNames"`
		IsUnique       bool     `json:"isUnique"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.AddSingleDataModelIndex(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName, reqBody.FieldNames, reqBody.IsUnique)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) DeleteSingleDataModelIndex(c *fiber.Ctx) error {
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		IndexName      string `json:"indexName"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	responseData, err := queryController.DeleteSingleDataModelIndex(reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    responseData,
	})
}

func (QueryHandlers) SaveDBQuery(c *fiber.Ctx) error {
	dbConnId := c.Params("dbConnId")
	var createBody struct {
		Name    string `json:"name"`
		Query   string `json:"query"`
		QueryID string `json:"queryId"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	analytics.SendSavedQueryEvent()
	queryObj, err := queryController.SaveDBQuery(dbConnId, createBody.Name, createBody.Query, createBody.QueryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildDBQueryView(queryObj),
	})
}

func (QueryHandlers) DeleteDBQuery(c *fiber.Ctx) error {
	queryID := c.Params("queryId")
	err := queryController.DeleteDBQuery(queryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}

func (QueryHandlers) GetDBQueriesInDBConnection(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	dbQueries, err := queryController.GetDBQueriesInDBConnection(dbConnID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	dbQueryViews := []views.DBQueryView{}
	for _, dbQuery := range dbQueries {
		dbQueryViews = append(dbQueryViews, *views.BuildDBQueryView(dbQuery))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    dbQueryViews,
	})
}

func (QueryHandlers) GetSingleDBQuery(c *fiber.Ctx) error {
	queryID := c.Params("queryId")
	dbQuery, err := queryController.GetSingleDBQuery(queryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildDBQueryView(dbQuery),
	})
}

func (QueryHandlers) GetQueryHistoryInDBConnection(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	beforeInt, err := strconv.ParseInt(c.Query("before"), 10, 64)
	var before time.Time
	if err != nil {
		before = time.Now()
	} else {
		before = utils.UnixNanoToTime(beforeInt)
	}
	dbQueryLogs, next, err := queryController.GetQueryHistoryInDBConnection(dbConnID, before)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	dbQueryLogViews := []views.DBQueryLogView{}
	for _, dbQueryLog := range dbQueryLogs {
		dbQueryLogViews = append(dbQueryLogViews, *views.BuildDBQueryLogView(dbQueryLog))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"list": dbQueryLogViews,
			"next": next,
		},
	})
}
