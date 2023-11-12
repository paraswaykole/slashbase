package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	commonviews "github.com/slashbaseide/slashbase/internal/common/views"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
	"github.com/slashbaseide/slashbase/internal/server/views"
)

type QueryHandlers struct{}

var queryController controllers.QueryController

func (QueryHandlers) RunQuery(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var runBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	if err := c.BodyParser(&runBody); err != nil {
		return fiber.ErrBadRequest
	}
	analytics.SendRunQueryEvent()
	response, err := queryController.RunQuery(authUser, runBody.DBConnectionID, runBody.Query)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
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
	var query struct {
		Filter []string `query:"filter"`
		Sort   []string `query:"sort"`
	}
	c.QueryParser(&query)
	analytics.SendLowCodeDataViewEvent()
	responsedata, err := queryController.GetData(authUser, authUserProjectIds, dbConnId, schema, name, fetchCount, limit, offset, query.Filter, query.Sort)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	dbConnId := c.Params("dbConnId")
	dataModels, err := queryController.GetDataModels(authUser, authUserProjectIds, dbConnId)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	dbConnId := c.Params("dbConnId")
	schema := c.Query("schema")
	name := c.Query("name")
	analytics.SendLowCodeModelViewEvent()
	data, err := queryController.GetSingleDataModel(authUser, authUserProjectIds, dbConnId, schema, name)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
		DataType       string `json:"dataType"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.AddSingleDataModelField(authUser, authUserProjectIds, reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName, reqBody.DataType)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		FieldName      string `json:"fieldName"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.DeleteSingleDataModelField(authUser, authUserProjectIds, reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.FieldName)
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
	authUser := middlewares.GetAuthUser(c)
	dbConnId := c.Params("dbConnId")
	var addBody struct {
		Schema string                 `json:"schema"`
		Name   string                 `json:"name"`
		Data   map[string]interface{} `json:"data"`
	}
	if err := c.BodyParser(&addBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.AddData(authUser, dbConnId, addBody.Schema, addBody.Name, addBody.Data)
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
	authUser := middlewares.GetAuthUser(c)
	dbConnId := c.Params("dbConnId")
	var deleteBody struct {
		Schema string   `json:"schema"`
		Name   string   `json:"name"`
		IDs    []string `json:"ids"` // ctid for postgres, _id for mongo
	}
	if err := c.BodyParser(&deleteBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.DeleteData(authUser, dbConnId, deleteBody.Schema, deleteBody.Name, deleteBody.IDs)
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
	authUser := middlewares.GetAuthUser(c)
	dbConnId := c.Params("dbConnId")
	var updateBody struct {
		Schema     string `json:"schema"`
		Name       string `json:"name"`
		ID         string `json:"id"`
		ColumnName string `json:"columnName"`
		Value      string `json:"value"`
	}
	if err := c.BodyParser(&updateBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.UpdateSingleData(authUser, dbConnId, updateBody.Schema, updateBody.Name, updateBody.ID, updateBody.ColumnName, updateBody.Value)
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
	authUser := middlewares.GetAuthUser(c)
	var reqBody struct {
		DBConnectionID string   `json:"dbConnectionId"`
		Schema         string   `json:"schema"`
		Name           string   `json:"name"`
		IndexName      string   `json:"indexName"`
		FieldNames     []string `json:"fieldNames"`
		IsUnique       bool     `json:"isUnique"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.AddSingleDataModelIndex(authUser, reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName, reqBody.FieldNames, reqBody.IsUnique)
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
	authUser := middlewares.GetAuthUser(c)
	var reqBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Schema         string `json:"schema"`
		Name           string `json:"name"`
		IndexName      string `json:"indexName"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	responseData, err := queryController.DeleteSingleDataModelIndex(authUser, reqBody.DBConnectionID, reqBody.Schema, reqBody.Name, reqBody.IndexName)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	dbConnId := c.Params("dbConnId")
	var createBody struct {
		Name    string `json:"name"`
		Query   string `json:"query"`
		QueryID string `json:"queryId"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return fiber.ErrBadRequest
	}
	analytics.SendSavedQueryEvent()
	queryObj, err := queryController.SaveDBQuery(authUser, authUserProjectIds, dbConnId, createBody.Name, createBody.Query, createBody.QueryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    commonviews.BuildDBQueryView(queryObj),
	})
}

func (QueryHandlers) DeleteDBQuery(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	queryID := c.Params("queryId")
	err := queryController.DeleteDBQuery(authUser, authUserProjectIds, queryID)
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
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	dbConnID := c.Params("dbConnId")
	dbQueries, err := queryController.GetDBQueriesInDBConnection(authUserProjectIds, dbConnID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	dbQueryViews := []commonviews.DBQueryView{}
	for _, dbQuery := range dbQueries {
		dbQueryViews = append(dbQueryViews, *commonviews.BuildDBQueryView(dbQuery))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    dbQueryViews,
	})
}

func (QueryHandlers) GetSingleDBQuery(c *fiber.Ctx) error {
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	queryID := c.Params("queryId")
	dbQuery, err := queryController.GetSingleDBQuery(authUserProjectIds, queryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    commonviews.BuildDBQueryView(dbQuery),
	})
}

func (QueryHandlers) GetQueryHistoryInDBConnection(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)
	dbConnID := c.Params("dbConnId")
	beforeInt, err := strconv.ParseInt(c.Query("before"), 10, 64)
	var before time.Time
	if err != nil {
		before = time.Now()
	} else {
		before = utils.UnixNanoToTime(beforeInt)
	}
	dbQueryLogs, next, err := queryController.GetQueryHistoryInDBConnection(authUser, authUserProjectIds, dbConnID, before)
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
