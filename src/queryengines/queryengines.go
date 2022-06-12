package queryengines

import (
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/queryengines/pgqueryengine"
)

var postgresQueryEngine *pgqueryengine.PostgresQueryEngine

func InitQueryEngines() {
	postgresQueryEngine = pgqueryengine.InitPostgresQueryEngine()
}

func RunQuery(user *models.User, dbConn *models.DBConnection, query string, userRole string) (map[string]interface{}, error) {
	return postgresQueryEngine.RunQuery(user, dbConn, query, true)
}

func TestConnection(user *models.User, dbConn *models.DBConnection) bool {
	return postgresQueryEngine.TestConnection(user, dbConn)
}

func GetDataModels(user *models.User, dbConn *models.DBConnection) ([]*DBDataModel, error) {
	data, err := postgresQueryEngine.GetDataModels(user, dbConn)
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

	dataModels := DBDataModel{
		SchemaName:  schemaName,
		Name:        name,
		Fields:      allFields,
		Constraints: allConstraints,
	}
	return &dataModels, nil
}

func GetData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string) (map[string]interface{}, error) {
	return postgresQueryEngine.GetData(user, dbConn, schemaName, name, limit, offset, fetchCount, filter, sort)
}

func UpdateSingleData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, ctid string, columnName, value string) (map[string]interface{}, error) {
	return postgresQueryEngine.UpdateSingleData(user, dbConn, schemaName, name, ctid, columnName, value)
}

func AddData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, data map[string]interface{}) (map[string]interface{}, error) {
	return postgresQueryEngine.AddData(user, dbConn, schemaName, name, data)
}

func DeleteData(user *models.User, dbConn *models.DBConnection, schemaName string, name string, ctids []string) (map[string]interface{}, error) {
	return postgresQueryEngine.DeleteData(user, dbConn, schemaName, name, ctids)
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
}
