package queryengines

import (
	"errors"

	"slashbase.com/backend/models"
)

var postgresQueryEngine PostgresQueryEngine

func RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	return postgresQueryEngine.RunQuery(dbConn, query)
}

func GetTables(dbConn *models.DBConnection) (map[string]interface{}, error) {
	data, err := postgresQueryEngine.GetTables(dbConn)
	if err != nil {
		return data, err
	}
	if data["success"] == false {
		return map[string]interface{}{}, errors.New(data["error"].(string))
	}
	return data["data"].(map[string]interface{}), nil
}
