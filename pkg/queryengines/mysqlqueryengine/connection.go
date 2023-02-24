package mysqlqueryengine

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/slashbaseide/slashbase/pkg/queryengines/utils"
)

type mysqlInstance struct {
	mysqlInstance *sql.DB
	LastUsed      time.Time
}

func (myEngine *MysqlQueryEngine) getConnection(dbConnectionId, host string, port uint16, database, user, password string) (c *sql.DB, err error) {
	if conn, exists := myEngine.openConnections[dbConnectionId]; exists {
		myEngine.mutex.Lock()
		myEngine.openConnections[dbConnectionId] = mysqlInstance{
			mysqlInstance: conn.mysqlInstance,
			LastUsed:      time.Now(),
		}
		myEngine.mutex.Unlock()
		return conn.mysqlInstance, nil
	}
	err = utils.CheckTcpConnection(host, strconv.Itoa(int(port)))
	if err != nil {
		return
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, strconv.Itoa(int(port)), database)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if dbConnectionId != "" {
		myEngine.mutex.Lock()
		myEngine.openConnections[dbConnectionId] = mysqlInstance{
			mysqlInstance: db,
			LastUsed:      time.Now(),
		}
		myEngine.mutex.Unlock()
	}
	return db, err
}

func (myEngine *MysqlQueryEngine) RemoveUnusedConnections() {
	for dbConnID, instance := range myEngine.openConnections {
		now := time.Now()
		diff := now.Sub(instance.LastUsed)
		if diff.Minutes() > 20 {
			myEngine.mutex.Lock()
			delete(myEngine.openConnections, dbConnID)
			myEngine.mutex.Unlock()
			go instance.mysqlInstance.Close()
		}
	}
}
