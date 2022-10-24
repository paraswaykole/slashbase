package queryengines

import (
	"errors"

	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/pkg/queryengines/mongoqueryengine"
	"slashbase.com/backend/pkg/queryengines/pgqueryengine"
)

var postgresQueryEngine *pgqueryengine.PostgresQueryEngine
var mongoQueryEngine *mongoqueryengine.MongoQueryEngine

func InitQueryEngines() {
	postgresQueryEngine = pgqueryengine.InitPostgresQueryEngine()
	mongoQueryEngine = mongoqueryengine.InitMongoQueryEngine()
}

func RunQuery(user *models.User, dbConn *models.DBConnection, query string, userRole string) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.RunQuery(user, dbConn, query, true)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.RunQuery(user, dbConn, query, true)
	}
	return nil, errors.New("invalid db type")
}

func TestConnection(user *models.User, dbConn *models.DBConnection) bool {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.TestConnection(user, dbConn)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.TestConnection(user, dbConn)
	}
	return false
}

func GetDataModels(user *models.User, dbConn *models.DBConnection) ([]*DBDataModel, error) {
	var err error
	var data []map[string]interface{}
	if dbConn.Type == models.DBTYPE_POSTGRES {
		data, err = postgresQueryEngine.GetDataModels(user, dbConn)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		data, err = mongoQueryEngine.GetDataModels(user, dbConn)
	}
	if err != nil {
		return nil, err
	}
	dataModels := []*DBDataModel{}
	for _, table := range data {
		view := BuildDBDataModel(dbConn, table)
		if view != nil {
			dataModels = append(dataModels, view)
		}
	}
	return dataModels, nil
}

func GetSingleDataModel(user *models.User, dbConn *models.DBConnection, schemaName string, name string) (*DBDataModel, error) {
	fieldsData, err := postgresQueryEngine.GetSingleDataModelFields(user, dbConn, schemaName, name)
	if err != nil {
		return nil, err
	}
	constraintsData, err := postgresQueryEngine.GetSingleDataModelConstraints(user, dbConn, schemaName, name)
	if err != nil {
		return nil, err
	}
	indexesData, err := postgresQueryEngine.GetSingleDataModelIndexes(user, dbConn, schemaName, name)
	if err != nil {
		return nil, err
	}
	allFields := []DBDataModelField{}
	for _, field := range fieldsData {
		fieldView := BuildDBDataModelField(dbConn, field)
		if fieldView != nil {
			allFields = append(allFields, *fieldView)
		}
	}
	allConstraints := []DBDataModelConstaint{}
	for _, constraint := range constraintsData {
		constraintView := BuildDBDataModelConstraint(dbConn, constraint)
		if constraintView != nil {
			allConstraints = append(allConstraints, *constraintView)
		}
	}
	allIndexes := []DBDataModelIndex{}
	for _, index := range indexesData {
		indexView := BuildDBDataModelIndex(dbConn, index)
		if indexView != nil {
			allIndexes = append(allIndexes, *indexView)
		}
	}

	dataModels := DBDataModel{
		SchemaName:  schemaName,
		Name:        name,
		Fields:      allFields,
		Constraints: allConstraints,
		Indexes:     allIndexes,
	}
	return &dataModels, nil
}

func GetData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.GetData(user, dbConn, schemaName, name, limit, offset, fetchCount, filter, sort)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.GetData(user, dbConn, name, limit, offset, fetchCount, filter, sort)
	}
	return nil, errors.New("invalid db type")
}

func UpdateSingleData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, ctid string, columnName, value string) (map[string]interface{}, error) {
	return postgresQueryEngine.UpdateSingleData(user, dbConn, schemaName, name, ctid, columnName, value)
}

func AddData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, data map[string]interface{}) (*AddDataResponse, error) {
	var result map[string]interface{}
	var err error
	if dbConn.Type == models.DBTYPE_POSTGRES {
		result, err = postgresQueryEngine.AddData(user, dbConn, schemaName, name, data)
		if err != nil {
			return nil, err
		}
	} else if dbConn.Type == models.DBTYPE_MONGO {
		result, err = mongoQueryEngine.AddData(user, dbConn, schemaName, name, data)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid db type")
	}
	return BuildAddDataResponse(dbConn, result), nil
}

// DeleteData function to delete multiple rows in the database
// ids is a list of unique row ids: ctid for postgres, _id for mongo
func DeleteData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, ids []string) (map[string]interface{}, error) {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		return postgresQueryEngine.DeleteData(user, dbConn, schemaName, name, ids)
	} else if dbConn.Type == models.DBTYPE_MONGO {
		return mongoQueryEngine.DeleteData(user, dbConn, name, ids)
	} else {
		return nil, errors.New("invalid db type")
	}
}

func CheckCreateRolePermissions(user *models.User, dbConn *models.DBConnection) bool {
	return postgresQueryEngine.CheckCreateRolePermissions(user, dbConn)
}

func CreateRoleLogin(user *models.User, dbConn *models.DBConnection, dbUser *models.DBConnectionUser) error {
	return postgresQueryEngine.CreateRoleLogin(user, dbConn, dbUser)
}

func DeleteRoleLogin(user *models.User, dbConn *models.DBConnection, dbUser *models.DBConnectionUser) error {
	return postgresQueryEngine.DeleteRoleLogin(user, dbConn, dbUser)
}

func RemoveUnusedConnections() {
	postgresQueryEngine.RemoveUnusedConnections()
	mongoQueryEngine.RemoveUnusedConnections()
}
