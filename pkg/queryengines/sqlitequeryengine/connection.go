package sqlitequeryengine

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteInstance struct {
	sqliteInstance *sql.DB
	LastUsed       time.Time
}

func (liteEngine *SqliteQueryEngine) getConnection(dbConnectionId, host string) (c *sql.DB, err error) {
	if conn, exists := liteEngine.openConnections[dbConnectionId]; exists {
		liteEngine.mutex.Lock()
		liteEngine.openConnections[dbConnectionId] = sqliteInstance{
			sqliteInstance: conn.sqliteInstance,
			LastUsed:       time.Now(),
		}
		liteEngine.mutex.Unlock()
		return conn.sqliteInstance, nil
	}
	db, err := sql.Open("sqlite3", host)
	if err != nil {
		return nil, err
	}
	if dbConnectionId != "" {
		liteEngine.mutex.Lock()
		liteEngine.openConnections[dbConnectionId] = sqliteInstance{
			sqliteInstance: db,
			LastUsed:       time.Now(),
		}
		liteEngine.mutex.Unlock()
	}
	return db, err
}

func (liteEngine *SqliteQueryEngine) RemoveUnusedConnections() {
	for dbConnID, instance := range liteEngine.openConnections {
		now := time.Now()
		diff := now.Sub(instance.LastUsed)
		if diff.Minutes() > 20 {
			liteEngine.mutex.Lock()
			delete(liteEngine.openConnections, dbConnID)
			liteEngine.mutex.Unlock()
			go instance.sqliteInstance.Close()
		}
	}
}
