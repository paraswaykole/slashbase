package queryengines

import (
	"slashbase.com/backend/models"
	"slashbase.com/backend/queryengines/pgqueryengine"
)

var postgresQueryEngine pgqueryengine.PostgresQueryEngine

func RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	return postgresQueryEngine.RunQuery(dbConn, query)
}

func GetDataModels(dbConn *models.DBConnection) (map[string]interface{}, error) {
	data, err := postgresQueryEngine.GetDataModels(dbConn)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetData(dbConn *models.DBConnection, schemaName string, name string, limit int, offset int64) (map[string]interface{}, error) {
	return postgresQueryEngine.GetData(dbConn, schemaName, name, limit, offset)
}

func RemoveUnusedConnections() {
	postgresQueryEngine.RemoveUnusedConnections()
}
