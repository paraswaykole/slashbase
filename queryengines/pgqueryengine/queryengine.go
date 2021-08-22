package pgqueryengine

import (
	"fmt"
	"strconv"
	"strings"

	"slashbase.com/backend/models"
	"slashbase.com/backend/models/sbsql"
	"slashbase.com/backend/queryengines/pgqueryengine/pgxutils"
	"slashbase.com/backend/sshtunnel"
)

type PostgresQueryEngine struct {
	openConnections map[string]pgxConnInstance
}

func (pgqe PostgresQueryEngine) RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	port, err := strconv.Atoi(string(dbConn.DBPort))
	if dbConn.UseSSH != models.DBUSESSH_NONE {
		sshTun := sshtunnel.GetSSHTunnel(dbConn.ID, dbConn.UseSSH,
			string(dbConn.SSHHost), port, string(dbConn.SSHUser),
			string(dbConn.SSHPassword), string(dbConn.SSHKeyFile),
		)
		dbConn.DBHost = "localhost"
		dbConn.DBPort = sbsql.CryptedData(fmt.Sprintf("%d", sshTun.GetLocalEndpoint().Port))
	}
	port, err = strconv.Atoi(string(dbConn.DBPort))
	conn, err := pgqe.getConnection(dbConn.ID, string(dbConn.DBHost), uint16(port), string(dbConn.DBName), string(dbConn.DBUser), string(dbConn.DBPassword))
	if err != nil {
		return nil, err
	}

	filteredQuery := strings.TrimSpace(strings.ToLower(query))
	if strings.HasPrefix(filteredQuery, "select") {
		rows, err := conn.Query(query)
		if err != nil {
			return nil, err
		}
		columns, rowsData := pgxutils.PgSqlRowsToJson(rows)
		return map[string]interface{}{
			"columns": columns,
			"rows":    rowsData,
		}, nil
	} else {
		cmdTag, err := conn.Exec(query)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message": cmdTag,
		}, nil
	}
}

func (pgqe PostgresQueryEngine) GetDataModels(dbConn *models.DBConnection) (map[string]interface{}, error) {
	return pgqe.RunQuery(dbConn, "SELECT * FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")
}

func (pgqe PostgresQueryEngine) GetData(dbConn *models.DBConnection, schema string, name string, limit int, offset int64, fetchCount bool) (map[string]interface{}, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s"."%s" LIMIT %d OFFSET %d;`, schema, name, limit, offset)
	data, err := pgqe.RunQuery(dbConn, query)
	if err != nil {
		return nil, err
	}
	if fetchCount {
		countQuery := fmt.Sprintf(`SELECT count(*) FROM "%s"."%s";`, schema, name)
		countData, err := pgqe.RunQuery(dbConn, countQuery)
		if err != nil {
			return nil, err
		}
		data["count"] = countData["rows"].([]map[string]interface{})[0]["count"]
	}
	return data, err
}
