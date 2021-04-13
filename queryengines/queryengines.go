package queryengines

import "slashbase.com/backend/models"

var postgresQueryEngine PostgresQueryEngine

func RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	return postgresQueryEngine.RunQuery(dbConn, query)
}
