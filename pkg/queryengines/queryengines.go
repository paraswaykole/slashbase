package queryengines

import (
	"errors"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/mongoqueryengine"
	"github.com/slashbaseide/slashbase/pkg/queryengines/mysqlqueryengine"
	"github.com/slashbaseide/slashbase/pkg/queryengines/pgqueryengine"
)

var postgresQueryEngine *pgqueryengine.PostgresQueryEngine
var mongoQueryEngine *mongoqueryengine.MongoQueryEngine
var mysqlQueryEngine *mysqlqueryengine.MysqlQueryEngine

func Init() {
	postgresQueryEngine = pgqueryengine.InitPostgresQueryEngine()
	mongoQueryEngine = mongoqueryengine.InitMongoQueryEngine()
	mysqlQueryEngine = mysqlqueryengine.InitMysqlQueryEngine()
}

func RunQuery(dbConn *models.DBConnection, query string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.RunQuery(dbConn, query, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.RunQuery(dbConn, query, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.RunQuery(dbConn, query, config)
	}
	return nil, errors.New("invalid db type")
}

func TestConnection(dbConn *models.DBConnection, config *models.QueryConfig) error {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.TestConnection(dbConn, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.TestConnection(dbConn, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.TestConnection(dbConn, config)
	}
	return errors.New("invalid db type")
}

func GetDataModels(dbConn *models.DBConnection, config *models.QueryConfig) ([]*models.DBDataModel, error) {
	var err error
	var data []map[string]interface{}
	if dbConn.Type == models.DBTYPE_POSTGRES {
		data, err = postgresQueryEngine.GetDataModels(dbConn, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		data, err = mongoQueryEngine.GetDataModels(dbConn, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		data, err = mysqlQueryEngine.GetDataModels(dbConn, config)
	}
	if err != nil {
		return nil, err
	}
	dataModels := []*models.DBDataModel{}
	for _, table := range data {
		view := models.BuildDBDataModel(dbConn, table)
		if view != nil {
			dataModels = append(dataModels, view)
		}
	}
	return dataModels, nil
}

func GetSingleDataModel(dbConn *models.DBConnection, schemaName string, name string, config *models.QueryConfig) (*models.DBDataModel, error) {
	var dataModel models.DBDataModel
	if dbConn.Type == models.DBTYPE_POSTGRES {
		fieldsData, err := postgresQueryEngine.GetSingleDataModelFields(dbConn, schemaName, name, config)
		if err != nil {
			return nil, err
		}
		indexesData, err := postgresQueryEngine.GetSingleDataModelIndexes(dbConn, schemaName, name, config)
		if err != nil {
			return nil, err
		}
		allFields := []models.DBDataModelField{}
		for _, field := range fieldsData {
			fieldView := models.BuildDBDataModelField(dbConn, field)
			if fieldView != nil {
				allFields = append(allFields, *fieldView)
			}
		}
		allIndexes := []models.DBDataModelIndex{}
		for _, index := range indexesData {
			indexView := models.BuildDBDataModelIndex(dbConn, index)
			if indexView != nil {
				allIndexes = append(allIndexes, *indexView)
			}
		}
		dataModel = models.DBDataModel{
			SchemaName: schemaName,
			Name:       name,
			Fields:     allFields,
			Indexes:    allIndexes,
		}
	} else if dbConn.Type == models.DBTYPE_MONGO {
		fieldsData, err := mongoQueryEngine.GetSingleDataModelFields(dbConn, name, config)
		if err != nil {
			return nil, err
		}
		indexesData, err := mongoQueryEngine.GetSingleDataModelIndexes(dbConn, name, config)
		if err != nil {
			return nil, err
		}
		allFields := []models.DBDataModelField{}
		for _, field := range fieldsData {
			fieldView := models.BuildDBDataModelField(dbConn, field)
			if fieldView != nil {
				allFields = append(allFields, *fieldView)
			}
		}
		allIndexes := []models.DBDataModelIndex{}
		for _, index := range indexesData {
			indexView := models.BuildDBDataModelIndex(dbConn, index)
			if indexView != nil {
				allIndexes = append(allIndexes, *indexView)
			}
		}
		dataModel = models.DBDataModel{
			Name:    name,
			Fields:  allFields,
			Indexes: allIndexes,
		}
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		fieldsData, err := mysqlQueryEngine.GetSingleDataModelFields(dbConn, name, config)
		if err != nil {
			return nil, err
		}
		indexesData, err := mysqlQueryEngine.GetSingleDataModelIndexes(dbConn, name, config)
		if err != nil {
			return nil, err
		}
		allFields := []models.DBDataModelField{}
		for _, field := range fieldsData {
			fieldView := models.BuildDBDataModelField(dbConn, field)
			if fieldView != nil {
				allFields = append(allFields, *fieldView)
			}
		}
		allIndexes := []models.DBDataModelIndex{}
		for _, index := range indexesData {
			indexView := models.BuildDBDataModelIndex(dbConn, index)
			if indexView != nil {
				allIndexes = append(allIndexes, *indexView)
			}
		}
		dataModel = models.DBDataModel{
			Name:    name,
			Fields:  allFields,
			Indexes: allIndexes,
		}
	}
	return &dataModel, nil
}

func AddSingleDataModelField(dbConn *models.DBConnection, schemaName string, name string, fieldName, datatype string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.AddSingleDataModelColumn(dbConn, schemaName, name, fieldName, datatype, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.AddSingleDataModelKey(dbConn, schemaName, name, fieldName, datatype)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.AddSingleDataModelColumn(dbConn, name, fieldName, datatype, config)
	}
	return nil, errors.New("invalid db type")
}

func DeleteSingleDataModelField(dbConn *models.DBConnection, schemaName string, name string, fieldName string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.DeleteSingleDataModelColumn(dbConn, schemaName, name, fieldName, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.DeleteSingleDataModelKey(dbConn, schemaName, name, fieldName, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.DeleteSingleDataModelColumn(dbConn, name, fieldName, config)
	}
	return nil, errors.New("invalid db type")
}

func GetData(dbConn *models.DBConnection, schemaName string, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.GetData(dbConn, schemaName, name, limit, offset, fetchCount, filter, sort, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.GetData(dbConn, name, limit, offset, fetchCount, filter, sort, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.GetData(dbConn, name, limit, offset, fetchCount, filter, sort, config)
	}
	return nil, errors.New("invalid db type")
}

// UpdateSingleData function to update single data row in the database
// id is a unique row ids: ctid for postgres, _id for mongo
func UpdateSingleData(dbConn *models.DBConnection, schemaName string, name string, id string, columnName, value string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.UpdateSingleData(dbConn, schemaName, name, id, columnName, value, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.UpdateSingleData(dbConn, name, id, value, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.UpdateSingleData(dbConn, name, id, columnName, value, config)
	}
	return nil, errors.New("invalid db type")
}

func AddData(dbConn *models.DBConnection, schemaName string, name string, data map[string]interface{}, config *models.QueryConfig) (*models.AddDataResponse, error) {
	var result map[string]interface{}
	var err error
	if dbConn.Type == models.DBTYPE_POSTGRES {
		result, err = postgresQueryEngine.AddData(dbConn, schemaName, name, data, config)
		if err != nil {
			return nil, err
		}
	} else if dbConn.Type == models.DBTYPE_MONGO {
		result, err = mongoQueryEngine.AddData(dbConn, schemaName, name, data, config)
		if err != nil {
			return nil, err
		}
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		result, err = mysqlQueryEngine.AddData(dbConn, name, data, config)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid db type")
	}
	return models.BuildAddDataResponse(dbConn, result), nil
}

// DeleteData function to delete multiple rows in the database
// ids is a list of unique row ids: ctid for postgres, _id for mongo
func DeleteData(dbConn *models.DBConnection, schemaName string, name string, ids []string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.DeleteData(dbConn, schemaName, name, ids, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.DeleteData(dbConn, name, ids, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.DeleteData(dbConn, name, ids, config)
	} else {
		return nil, errors.New("invalid db type")
	}
}

func AddSingleDataModelIndex(dbConn *models.DBConnection, schemaName, name, indexName string, fieldNames []string, isUnique bool, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.AddSingleDataModelIndex(dbConn, schemaName, name, indexName, fieldNames, isUnique, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.AddSingleDataModelIndex(dbConn, name, indexName, fieldNames, isUnique, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.AddSingleDataModelIndex(dbConn, name, indexName, fieldNames, isUnique, config)
	} else {
		return nil, errors.New("invalid db type")
	}
}

func DeleteSingleDataModelIndex(dbConn *models.DBConnection, schemaName, name, indexName string, config *models.QueryConfig) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.DeleteSingleDataModelIndex(dbConn, indexName, config)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.DeleteSingleDataModelIndex(dbConn, name, indexName, config)
	} else if dbConn.Type == models.DBTYPE_MYSQL {
		return mysqlQueryEngine.DeleteSingleDataModelIndex(dbConn, name, indexName, config)
	} else {
		return nil, errors.New("invalid db type")
	}
}

func RemoveUnusedConnections() {
	postgresQueryEngine.RemoveUnusedConnections()
	mongoQueryEngine.RemoveUnusedConnections()
	mysqlQueryEngine.RemoveUnusedConnections()
}
