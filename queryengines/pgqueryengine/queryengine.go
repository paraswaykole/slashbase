package pgqueryengine

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"slashbase.com/backend/models"
	"slashbase.com/backend/models/sbsql"
	"slashbase.com/backend/queryengines/pgqueryengine/pgxutils"
	"slashbase.com/backend/sshtunnel"
)

type PostgresQueryEngine struct {
	openConnections map[string]pgxConnPoolInstance
}

func InitPostgresQueryEngine() *PostgresQueryEngine {
	return &PostgresQueryEngine{
		openConnections: map[string]pgxConnPoolInstance{},
	}
}

func (pgqe *PostgresQueryEngine) RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	port, _ := strconv.Atoi(string(dbConn.DBPort))
	if dbConn.UseSSH != models.DBUSESSH_NONE {
		sshTun := sshtunnel.GetSSHTunnel(dbConn.ID, dbConn.UseSSH,
			string(dbConn.SSHHost), port, string(dbConn.SSHUser),
			string(dbConn.SSHPassword), string(dbConn.SSHKeyFile),
		)
		dbConn.DBHost = "localhost"
		dbConn.DBPort = sbsql.CryptedData(fmt.Sprintf("%d", sshTun.GetLocalEndpoint().Port))
	}
	port, _ = strconv.Atoi(string(dbConn.DBPort))
	conn, err := pgqe.getConnection(dbConn.ID, string(dbConn.DBHost), uint16(port), string(dbConn.DBName), string(dbConn.DBUser), string(dbConn.DBPassword))
	if err != nil {
		return nil, err
	}

	queryType := pgxutils.GetPSQLQueryType(query)
	if queryType == pgxutils.QUERY_READ {
		rows, err := conn.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}
		columns, rowsData := pgxutils.PgSqlRowsToJson(rows)
		return map[string]interface{}{
			"columns": columns,
			"rows":    rowsData,
		}, nil
	} else {
		cmdTag, err := conn.Exec(context.Background(), query)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message": cmdTag.String(),
		}, nil
	}
}

func (pgqe *PostgresQueryEngine) GetDataModels(dbConn *models.DBConnection) (map[string]interface{}, error) {
	return pgqe.RunQuery(dbConn, "SELECT * FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';")
}

func (pgqe *PostgresQueryEngine) GetData(dbConn *models.DBConnection, schema string, name string, limit int, offset int64, fetchCount bool) (map[string]interface{}, error) {
	query := fmt.Sprintf(`SELECT ctid, * FROM "%s"."%s" LIMIT %d OFFSET %d;`, schema, name, limit, offset)
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

func (pgqe *PostgresQueryEngine) UpdateSingleData(dbConn *models.DBConnection, schema string, name string, ctid string, columnName string, value string) (map[string]interface{}, error) {
	query := fmt.Sprintf(`UPDATE "%s"."%s" SET "%s" = '%s' WHERE ctid = '%s' RETURNING ctid;`, schema, name, columnName, value, ctid)
	data, err := pgqe.RunQuery(dbConn, query)
	if err != nil {
		return nil, err
	}
	ctID := data["rows"].([]map[string]interface{})[0]["ctid"]
	data = map[string]interface{}{
		"ctid": ctID,
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) AddData(dbConn *models.DBConnection, schema string, name string, data map[string]interface{}) (map[string]interface{}, error) {
	keys := []string{}
	values := []string{}
	for key, value := range data {
		keys = append(keys, key)
		val := value.(string)
		values = append(values, val)
	}
	keysStr := strings.Join(keys, ", ")
	valuesStr := strings.Join(values, "','")
	query := fmt.Sprintf(`INSERT INTO "%s"."%s"(%s) VALUES('%s') RETURNING ctid;`, schema, name, keysStr, valuesStr)
	rData, err := pgqe.RunQuery(dbConn, query)
	if err != nil {
		return nil, err
	}
	ctID := rData["rows"].([]map[string]interface{})[0]["ctid"]
	rData = map[string]interface{}{
		"ctid": ctID,
	}
	return rData, err
}

func (pgqe *PostgresQueryEngine) DeleteData(dbConn *models.DBConnection, schema string, name string, ctids []string) (map[string]interface{}, error) {
	ctidsStr := strings.Join(ctids, "', '")
	query := fmt.Sprintf(`DELETE FROM "%s"."%s" WHERE ctid IN ('%s');`, schema, name, ctidsStr)
	return pgqe.RunQuery(dbConn, query)
}
