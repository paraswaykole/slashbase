package controllers

import (
	"errors"
	"time"

	"github.com/slashbaseide/slashbase/internal/common/config"
	commondao "github.com/slashbaseide/slashbase/internal/common/dao"
	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

type QueryController struct{}

func (QueryController) RunQuery(authUser *models.User, dbConnectionId, query string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnectionId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.RunQuery(dbConn.ToQEConnection(), query, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetData(authUser *models.User, authUserProjectIds *[]string,
	dbConnId, schema, name string, fetchCount bool, limit int, offset int64,
	filter, sort []string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.GetData(dbConn.ToQEConnection(), schema, name, limit, offset, fetchCount, filter, sort, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetDataModels(authUser *models.User, authUserProjectIds *[]string, dbConnId string) ([]*qemodels.DBDataModel, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	dataModels, err := queryengines.GetDataModels(dbConn.ToQEConnection(), getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (QueryController) GetSingleDataModel(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string) (*qemodels.DBDataModel, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.GetSingleDataModel(dbConn.ToQEConnection(), schema, name, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddSingleDataModelField(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string, fieldName, dataType string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}
	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.AddSingleDataModelField(dbConn.ToQEConnection(), schema, name, fieldName, dataType, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) DeleteSingleDataModelField(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	schema, name string, fieldName string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed to run query")
	}
	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.DeleteSingleDataModelField(dbConn.ToQEConnection(), schema, name, fieldName, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddData(authUser *models.User, dbConnId string,
	schema, name string, data map[string]interface{}) (*qemodels.AddDataResponse, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	resultData, err := queryengines.AddData(dbConn.ToQEConnection(), schema, name, data, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return resultData, nil
}

func (QueryController) DeleteData(authUser *models.User, dbConnId string,
	schema, name string, ids []string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.DeleteData(dbConn.ToQEConnection(), schema, name, ids, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) UpdateSingleData(authUser *models.User, dbConnId string,
	schema, name, id, columnName, columnValue string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.UpdateSingleData(dbConn.ToQEConnection(), schema, name, id, columnName, columnValue, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) AddSingleDataModelIndex(authUser *models.User, dbConnId string,
	schema, name string, indexName string, fieldNames []string, isUnique bool) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.AddSingleDataModelIndex(dbConn.ToQEConnection(), schema, name, indexName, fieldNames, isUnique, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) DeleteSingleDataModelIndex(authUser *models.User, dbConnId string,
	schema, name string, indexName string) (map[string]interface{}, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return nil, err
	}

	data, err := queryengines.DeleteSingleDataModelIndex(dbConn.ToQEConnection(), schema, name, indexName, getQueryConfigsForProjectMember(pm, dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) SaveDBQuery(authUser *models.User, authUserProjectIds *[]string, dbConnId string,
	name, query, queryId string) (*common.DBQuery, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed")
	}

	var queryObj *common.DBQuery
	if queryId == "" {
		queryObj = common.NewQuery(name, query, dbConn.ID)
		err = commondao.DBQuery.CreateQuery(queryObj)
	} else {
		queryObj, err = commondao.DBQuery.GetSingleDBQuery(queryId)
		if err != nil {
			return nil, errors.New("there was some problem")
		}
		queryObj.Name = name
		queryObj.Query = query
		err = commondao.DBQuery.UpdateDBQuery(queryId, &common.DBQuery{
			Name:  name,
			Query: query,
		})
	}

	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return queryObj, nil
}

func (QueryController) DeleteDBQuery(authUser *models.User, authUserProjectIds *[]string, queryId string) error {

	query, err := commondao.DBQuery.GetSingleDBQuery(queryId)
	if err != nil {
		return errors.New("there was some problem")
	}

	if !utils.ContainsString(*authUserProjectIds, query.DBConnection.ProjectID) {
		return errors.New("not allowed")
	}

	err = commondao.DBQuery.DeleteDBQuery(queryId)
	if err != nil {
		return errors.New("there was some problem")
	}
	return nil
}

func (QueryController) GetDBQueriesInDBConnection(authUserProjectIds *[]string, dbConnId string) ([]*common.DBQuery, error) {

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, errors.New("not allowed")
	}

	dbQueries, err := commondao.DBQuery.GetDBQueriesByDBConnId(dbConnId)
	if err != nil {
		return nil, err
	}
	return dbQueries, nil
}

func (QueryController) GetSingleDBQuery(authUserProjectIds *[]string, queryId string) (*common.DBQuery, error) {

	dbQuery, err := commondao.DBQuery.GetSingleDBQuery(queryId)
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

	dbConn, err := commondao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}
	if !utils.ContainsString(*authUserProjectIds, dbConn.ProjectID) {
		return nil, 0, errors.New("not allowed")
	}

	dbQueryLogs, err := dao.DBQueryLog.GetDBQueryLogsDBConnID(dbConnId, before)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}

	var next int64 = -1
	if len(dbQueryLogs) == config.PAGINATION_COUNT {
		next = dbQueryLogs[len(dbQueryLogs)-1].CreatedAt.UnixNano()
	}

	return dbQueryLogs, next, nil
}
