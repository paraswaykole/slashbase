package sqlitequeryengine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/sqlitequeryengine/sqliteutils"
)

type SqliteQueryEngine struct {
	openConnections map[string]sqliteInstance
	mutex           *sync.Mutex
}

func InitSqliteQueryEngine() *SqliteQueryEngine {
	return &SqliteQueryEngine{
		openConnections: map[string]sqliteInstance{},
		mutex:           &sync.Mutex{},
	}
}

func (slqe *SqliteQueryEngine) RunQuery(dbConn *models.DBConnection, query string, config *models.QueryConfig) (map[string]interface{}, error) {
	conn, err := slqe.getConnection(dbConn.ID, string(dbConn.DBHost))
	if err != nil {
		return nil, err
	}

	queryType, isReturningRows := sqliteutils.GetSQLiteQueryType(query)

	if queryType != sqliteutils.QUERY_READ && config.ReadOnly {
		return nil, errors.New("not allowed run this query")
	}

	if isReturningRows {
		rows, err := conn.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		columns, rowsData := sqliteutils.SqliteRowsToJson(rows)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"columns": columns,
			"rows":    rowsData,
		}, nil
	}
	result, err := conn.Exec(query)
	if err != nil {
		return nil, err
	}
	if config.CreateLogFn != nil {
		config.CreateLogFn(query)
	}
	rowsAffected, _ := result.RowsAffected()
	return map[string]interface{}{
		"message": fmt.Sprintf("%d rows affected", rowsAffected),
	}, nil
}
