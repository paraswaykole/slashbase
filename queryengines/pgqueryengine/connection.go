package pgqueryengine

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
)

type pgxConnInstance struct {
	pgxConnInstance *pgx.Conn
	LastUsed        time.Time
}

func (pxEngine PostgresQueryEngine) getConnection(dbConnectionId, host string, port uint16, database, user, password string) (c *pgx.Conn, err error) {
	if conn, exists := pxEngine.openConnections[dbConnectionId]; exists {
		pxEngine.openConnections[dbConnectionId] = pgxConnInstance{
			pgxConnInstance: conn.pgxConnInstance,
			LastUsed:        time.Now(),
		}
		return conn.pgxConnInstance, nil
	}
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	})
	if err != nil {
		if pgerr, ok := err.(pgx.PgError); ok {
			err = errors.New(fmt.Sprintf("Unable to connect to database: %v", pgerr.Message))
		}
		err = errors.New(fmt.Sprintf("Unable to connect to database: %v", err))
		return
	}
	if pxEngine.openConnections == nil {
		pxEngine.openConnections = map[string]pgxConnInstance{}
	}
	pxEngine.openConnections[dbConnectionId] = pgxConnInstance{
		pgxConnInstance: conn,
		LastUsed:        time.Now(),
	}
	return conn, err
}

func (pxEngine PostgresQueryEngine) RemoveUnusedConnections() {
	for {
		time.Sleep(time.Minute * time.Duration(5))
		for dbConnID, instance := range pxEngine.openConnections {
			now := time.Now()
			diff := now.Sub(instance.LastUsed)
			if diff.Minutes() > 20 {
				delete(pxEngine.openConnections, dbConnID)
				go instance.pgxConnInstance.Close()
			}
		}
	}
}
