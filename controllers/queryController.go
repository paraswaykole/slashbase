package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/queryengines"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type QueryController struct{}

var dbQueryDao daos.DBQueryDao

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

	authUserProjectMember, err := middlewares.GetAuthUserProjectMemberForProject(c, dbConn.ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	data, err := queryengines.RunQuery(dbConn, runBody.Query, authUserProjectMember.Role)
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dataModels,
	})
}

func (qc QueryController) GetSingleDataModel(c *gin.Context) {
	dbConnId := c.Param("dbConnId")

	schema := c.Query("schema")
	name := c.Query("name")

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

	data, err := queryengines.GetSingleDataModel(dbConn, schema, name)
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

	data, err := queryengines.AddData(dbConn, addBody.Schema, addBody.Name, addBody.Data)
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

	var deleteBody struct {
		Schema string   `json:"schema"`
		Name   string   `json:"name"`
		CTIDs  []string `json:"ctids"`
	}
	c.BindJSON(&deleteBody)

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

	data, err := queryengines.DeleteData(dbConn, deleteBody.Schema, deleteBody.Name, deleteBody.CTIDs)
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
