package mysqlqueryengine

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/mysqlqueryengine/mysqlutils"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
)

type MysqlQueryEngine struct {
	openConnections map[string]mysqlInstance
}

func InitMysqlQueryEngine() *MysqlQueryEngine {
	return &MysqlQueryEngine{
		openConnections: map[string]mysqlInstance{},
	}
}

func (mqe *MysqlQueryEngine) RunQuery(dbConn *models.DBConnection, query string, config *models.QueryConfig) (map[string]interface{}, error) {
	port, _ := strconv.Atoi(string(dbConn.DBPort))
	if dbConn.UseSSH != models.DBUSESSH_NONE {
		remoteHost := string(dbConn.DBHost)
		if remoteHost == "" {
			remoteHost = "localhost"
		}
		sshTun := sshtunnel.GetSSHTunnel(dbConn.ID, dbConn.UseSSH,
			string(dbConn.SSHHost), remoteHost, port, string(dbConn.SSHUser),
			string(dbConn.SSHPassword), string(dbConn.SSHKeyFile),
		)
		dbConn.DBHost = "localhost"
		dbConn.DBPort = fmt.Sprintf("%d", sshTun.GetLocalEndpoint().Port)
	}
	port, _ = strconv.Atoi(string(dbConn.DBPort))
	conn, err := mqe.getConnection(dbConn.ID, string(dbConn.DBHost), uint16(port), string(dbConn.DBName), string(dbConn.DBUser), string(dbConn.DBPassword))
	if err != nil {
		return nil, err
	}

	queryType, isReturningRows := mysqlutils.GetMySQLQueryType(query)

	if queryType != mysqlutils.QUERY_READ && config.ReadOnly {
		return nil, errors.New("not allowed run this query")
	}

	if isReturningRows {
		rows, err := conn.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		columns, rowsData := mysqlutils.MySqlRowsToJson(rows)
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
