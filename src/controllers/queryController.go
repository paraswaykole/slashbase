package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/queryengines"
	"slashbase.com/backend/src/utils"
	"slashbase.com/backend/src/views"
)

type QueryController struct{}

var dbQueryDao daos.DBQueryDao
var dbQueryLogDao daos.DBQueryLogDao

func (qc QueryController) RunQuery(c *gin.Context) {
	var runBody struct {
		DBConnectionID string `json:"dbConnectionId"`
		Query          string `json:"query"`
	}
	c.BindJSON(&runBody)
	authUser := middlewares.GetAuthUser(c)

	dbConn, err := dbConnDao.GetConnectableDBConnection(runBody.DBConnectionID, authUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}

	authUserProjectMember, err := middlewares.GetAuthUserProjectMemberForProject(c, dbConn.ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	data, err := queryengines.RunQuery(authUser, dbConn, runBody.Query, authUserProjectMember.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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
	filter, hasFilter := c.GetQueryArray("filter[]")
	if hasFilter && len(filter) < 2 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Invalid filter query",
		})
		return
	}
	sort, hasSort := c.GetQueryArray("sort[]")
	if hasSort && len(sort) != 2 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Invalid sort query",
		})
		return
	}
	authUser := middlewares.GetAuthUser(c)
	authUserProjects := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	data, err := queryengines.GetData(authUser, dbConn, schema, name, limit, offset, fetchCount, filter, sort)
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
	authUser := middlewares.GetAuthUser(c)
	authUserProjects := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	dataModels, err := queryengines.GetDataModels(authUser, dbConn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dataModels,
	})
}

func (qc QueryController) GetSingleDataModel(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	schema := c.Query("schema")
	name := c.Query("name")
	authUser := middlewares.GetAuthUser(c)
	authUserProjects := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	data, err := queryengines.GetSingleDataModel(authUser, dbConn, schema, name)
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

func (qc QueryController) AddData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	var addBody struct {
		Schema string                 `json:"schema"`
		Name   string                 `json:"name"`
		Data   map[string]interface{} `json:"data"`
	}
	c.BindJSON(&addBody)
	authUser := middlewares.GetAuthUser(c)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	data, err := queryengines.AddData(authUser, dbConn, addBody.Schema, addBody.Name, addBody.Data)
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

func (qc QueryController) DeleteData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	var deleteBody struct {
		Schema string   `json:"schema"`
		Name   string   `json:"name"`
		CTIDs  []string `json:"ctids"`
	}
	c.BindJSON(&deleteBody)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	data, err := queryengines.DeleteData(authUser, dbConn, deleteBody.Schema, deleteBody.Name, deleteBody.CTIDs)
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

func (qc QueryController) UpdateSingleData(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	var updateBody struct {
		Schema     string `json:"schema"`
		Name       string `json:"name"`
		CTID       string `json:"ctid"`
		ColumnName string `json:"columnName"`
		Value      string `json:"value"`
	}
	c.BindJSON(&updateBody)

	dbConn, err := dbConnDao.GetConnectableDBConnection(dbConnId, authUser.ID)
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

	data, err := queryengines.UpdateSingleData(authUser, dbConn, updateBody.Schema, updateBody.Name, updateBody.CTID, updateBody.ColumnName, updateBody.Value)
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

func (qc QueryController) SaveDBQuery(c *gin.Context) {
	dbConnId := c.Param("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	authUserProjects := middlewares.GetAuthUserProjectIds(c)
	var createBody struct {
		Name    string `json:"name"`
		Query   string `json:"query"`
		QueryID string `json:"queryId"`
	}
	c.BindJSON(&createBody)

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

	var queryObj *models.DBQuery
	if createBody.QueryID == "" {
		queryObj = models.NewQuery(authUser, createBody.Name, createBody.Query, dbConn.ID)
		err = dbQueryDao.CreateQuery(queryObj)
	} else {
		queryObj, err = dbQueryDao.GetSingleDBQuery(createBody.QueryID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		queryObj.Name = createBody.Name
		queryObj.Query = createBody.Query
		err = dbQueryDao.UpdateDBQuery(createBody.QueryID, &models.DBQuery{
			Name:  createBody.Name,
			Query: createBody.Query,
		})

	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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

func (qc QueryController) GetDBQueriesInDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)

	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	dbQueries, err := dbQueryDao.GetDBQueriesByDBConnId(dbConnID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
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

func (qc QueryController) GetSingleDBQuery(c *gin.Context) {
	queryID := c.Param("queryId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)

	dbQuery, err := dbQueryDao.GetSingleDBQuery(queryID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if !utils.ContainsString(*authUserProjectIds, dbQuery.DBConnection.ProjectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildDBQueryView(dbQuery),
	})
}

func (qc QueryController) GetQueryHistoryInDBConnection(c *gin.Context) {
	dbConnID := c.Param("dbConnId")
	authUserProjectIds := middlewares.GetAuthUserProjectIds(c)

	beforeInt, err := strconv.ParseInt(c.Query("before"), 10, 64)
	var before time.Time
	if err != nil {
		before = time.Now()
	} else {
		before = utils.UnixNanoToTime(beforeInt)
	}

	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}

	authUserProjectMember, err := middlewares.GetAuthUserProjectMemberForProject(c, dbConn.ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	dbQueryLogs, err := dbQueryLogDao.GetDBQueryLogsDBConnID(dbConnID, authUserProjectMember, before)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	dbQueryLogViews := []views.DBQueryLogView{}
	var next int64 = -1
	for i, dbQueryLog := range dbQueryLogs {
		dbQueryLogViews = append(dbQueryLogViews, *views.BuildDBQueryLogView(dbQueryLog))
		if i == config.PAGINATION_COUNT-1 {
			next = dbQueryLog.CreatedAt.UnixNano()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"list": dbQueryLogViews,
			"next": next,
		},
	})
}
