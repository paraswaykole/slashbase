package controllers

import (
	"errors"
	"time"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

type QueryController struct{}

func (QueryController) RunQuery(dbConnectionId, query string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.RunQuery(dbConn.ToQEConnection(), query, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetData(dbConnId, schema, name string, isFirstFetch bool, limit int, offset int64,
	filter, sort []string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.GetData(dbConn.ToQEConnection(), schema, name, limit, offset, isFirstFetch, filter, sort, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) GetDataModels(dbConnId string) ([]*qemodels.DBDataModel, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	dataModels, err := queryengines.GetDataModels(dbConn.ToQEConnection(), getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return dataModels, nil
}

func (QueryController) GetSingleDataModel(dbConnId string, schema, name string) (*qemodels.DBDataModel, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.GetSingleDataModel(dbConn.ToQEConnection(), schema, name, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddSingleDataModelField(dbConnId string, schema, name string, fieldName, dataType string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.AddSingleDataModelField(dbConn.ToQEConnection(), schema, name, fieldName, dataType, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) DeleteSingleDataModelField(dbConnId string,
	schema, name string, fieldName string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.DeleteSingleDataModelField(dbConn.ToQEConnection(), schema, name, fieldName, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) AddData(dbConnId string,
	schema, name string, data map[string]interface{}) (*qemodels.AddDataResponse, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	resultData, err := queryengines.AddData(dbConn.ToQEConnection(), schema, name, data, getQueryConfigs(dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return resultData, nil
}

func (QueryController) DeleteData(dbConnId string,
	schema, name string, ids []string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.DeleteData(dbConn.ToQEConnection(), schema, name, ids, getQueryConfigs(dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) UpdateSingleData(dbConnId string,
	schema, name, id, columnName, columnValue string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.UpdateSingleData(dbConn.ToQEConnection(), schema, name, id, columnName, columnValue, getQueryConfigs(dbConn))
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return data, nil
}

func (QueryController) AddSingleDataModelIndex(dbConnId string,
	schema, name string, indexName string, fieldNames []string, isUnique bool) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.AddSingleDataModelIndex(dbConn.ToQEConnection(), schema, name, indexName, fieldNames, isUnique, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) DeleteSingleDataModelIndex(dbConnId string,
	schema, name string, indexName string) (map[string]interface{}, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	data, err := queryengines.DeleteSingleDataModelIndex(dbConn.ToQEConnection(), schema, name, indexName, getQueryConfigs(dbConn))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (QueryController) SaveDBQuery(dbConnId string,
	name, query, queryId string) (*models.DBQuery, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	var queryObj *models.DBQuery
	if queryId == "" {
		queryObj = models.NewQuery(name, query, dbConn.ID)
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

func (QueryController) DeleteDBQuery(queryId string) error {

	query, err := dao.DBQuery.GetSingleDBQuery(queryId)
	if err != nil {
		return errors.New("there was some problem")
	}

	err = dao.DBQuery.DeleteDBQuery(query.ID)
	if err != nil {
		return errors.New("there was some problem")
	}
	return nil
}

func (QueryController) GetDBQueriesInDBConnection(dbConnId string) ([]*models.DBQuery, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	dbQueries, err := dao.DBQuery.GetDBQueriesByDBConnId(dbConn.ID)
	if err != nil {
		return nil, err
	}
	return dbQueries, nil
}

func (QueryController) GetSingleDBQuery(queryId string) (*models.DBQuery, error) {

	dbQuery, err := dao.DBQuery.GetSingleDBQuery(queryId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return dbQuery, nil
}

func (QueryController) GetQueryHistoryInDBConnection(dbConnId string, before time.Time) ([]*models.DBQueryLog, int64, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}

	dbQueryLogs, err := dao.DBQueryLog.GetDBQueryLogsDBConnID(dbConn.ID, before)
	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}

	var next int64 = -1
	if len(dbQueryLogs) == config.PAGINATION_COUNT {
		next = dbQueryLogs[len(dbQueryLogs)-1].CreatedAt.UnixNano()
	}

	return dbQueryLogs, next, nil
}
