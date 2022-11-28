package controllers

import (
	"errors"
	"time"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/pkg/queryengines"
)

type QueryController struct{}

func (QueryController) RunQuery(authUser *models.User, dbConnectionId, query string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.RunQuery(dbConn, query, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetData(authUser *models.User, authUserProjectIds *[]string,
	dbConnId, schema, name string, fetchCount bool, limit int, offset int64,
	filter, sort []string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.GetData(dbConn, schema, name, limit, offset, fetchCount, filter, sort, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetDataModels(authUser *models.User, authUserProjectIds *[]string, dbConnId string) ([]*queryengines.DBDataModel, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	dataModels, err := queryengines.GetDataModels(dbConn, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (QueryController) GetSingleDataModel(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string) (*queryengines.DBDataModel, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.GetSingleDataModel(dbConn, schema, name, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddSingleDataModelField(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string, fieldName, dataType string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}
	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.AddSingleDataModelField(dbConn, schema, name, fieldName, dataType, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) DeleteSingleDataModelField(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string, fieldName string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}
	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.DeleteSingleDataModelField(dbConn, schema, name, fieldName, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddData(authUser *models.User, dbConnId string,
	schema, name string, data map[string]interface{}) (*queryengines.AddDataResponse, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	resultData, err := queryengines.AddData(dbConn, schema, name, data, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return resultData, nil
}

func (QueryController) DeleteData(authUser *models.User, dbConnId string,
	schema, name string, ids []string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.DeleteData(dbConn, schema, name, ids, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) UpdateSingleData(authUser *models.User, dbConnId string,
	schema, name, id, columnName, columnValue string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.UpdateSingleData(dbConn, schema, name, id, columnName, columnValue, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) SaveDBQuery(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	name, query, queryId string) (*models.DBQuery, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed")
	}

	var queryObj *models.DBQuery
	if queryId == "" {
		queryObj = models.NewQuery(authUser, name, query, dbConn.ID)
		err = dao.DBQuery.CreateQuery(queryObj)
	} else {
		queryObj, err = dao.DBQuery.GetSingleDBQuery(queryId)
		if err != nil {
			return nil, errors.New("there was some problem")
		}
		queryObj.Name = name
		queryObj.Query = query
		err = dao.DBQuery.UpdateDBQuery(queryId, &models.DBQuery{
			Name:  name,
			Query: query,
		})
	}

	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return queryObj, nil
}

func (QueryController) GetDBQueriesInDBConnection(authUserProjectIds *[]string, dbConnId string) ([]*models.DBQuery, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed")
	}

	dbQueries, err := dao.DBQuery.GetDBQueriesByDBConnId(dbConnId)
	if err != nil {
		return nil, err
	}
	return dbQueries, nil
}

func (QueryController) GetSingleDBQuery(authUserProjectIds *[]string, queryId string) (*models.DBQuery, error) {

	dbQuery, err := dao.DBQuery.GetSingleDBQuery(queryId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if !utils.ContainsString(*authUserProjectIds, dbQuery.DBConnection.ProjectID) {
		return nil, errors.New("not allowed")
	}

	return dbQuery, nil
}

func (QueryController) GetQueryHistoryInDBConnection(authUser *models.User, authUserProjectIds *[]string,
	dbConnId string, before time.Time) ([]*models.DBQueryLog, int64, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, 0, errors.New("not allowed")
	}

	authUserProjectMember, err := getAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, 0, err
	}

	dbQueryLogs, err := dao.DBQueryLog.GetDBQueryLogsDBConnID(dbConnId, authUserProjectMember, before)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}

	var next int64 = -1
	if len(dbQueryLogs) == config.PAGINATION_COUNT {
		next = dbQueryLogs[len(dbQueryLogs)-1].CreatedAt.UnixNano()
	}

	return dbQueryLogs, next, nil
}
