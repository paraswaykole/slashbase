package pgqueryengine

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/pgqueryengine/pgxutils"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
)

type PostgresQueryEngine struct {
	openConnections map[string]pgxConnPoolInstance
	mutex           *sync.Mutex
}

func InitPostgresQueryEngine() *PostgresQueryEngine {
	return &PostgresQueryEngine{
		openConnections: map[string]pgxConnPoolInstance{},
		mutex:           &sync.Mutex{},
	}
}

func (pgqe *PostgresQueryEngine) RunQuery(dbConn *models.DBConnection, query string, config *models.QueryConfig) (map[string]interface{}, error) {
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
	conn, err := pgqe.getConnection(dbConn.ID, string(dbConn.DBHost), uint16(port), string(dbConn.DBName), string(dbConn.DBUser), string(dbConn.DBPassword))
	if err != nil {
		return nil, err
	}

	queryType, isReturningRows := pgxutils.GetPSQLQueryType(query)

	if queryType != pgxutils.QUERY_READ && config.ReadOnly {
		return nil, errors.New("not allowed run this query")
	}

	if isReturningRows {
		rows, err := conn.Query(context.Background(), query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		columns, rowsData := pgxutils.PgSqlRowsToJson(rows)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"columns": columns,
			"rows":    rowsData,
		}, nil
	}
	cmdTag, err := conn.Exec(context.Background(), query)
	if err != nil {
		return nil, err
	}
	if config.CreateLogFn != nil {
		config.CreateLogFn(query)
	}
	return map[string]interface{}{
		"message": cmdTag.String(),
	}, nil
}

func (pgqe *PostgresQueryEngine) TestConnection(dbConn *models.DBConnection, config *models.QueryConfig) error {
	query := "SELECT 1 AS test;"
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return err
	}
	test := data["rows"].([]map[string]interface{})[0]["0"].(int32)
	if test == 1 {
		return nil
	}
	return errors.New("connection test failed")
}

func (pgqe *PostgresQueryEngine) GetDataModels(dbConn *models.DBConnection, config *models.QueryConfig) ([]map[string]interface{}, error) {
	data, err := pgqe.RunQuery(dbConn, "SELECT tablename, schemaname FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema' ORDER BY tablename;", config)
	if err != nil {
		return nil, err
	}
	rdata := data["rows"].([]map[string]interface{})
	return rdata, nil
}

func (pgqe *PostgresQueryEngine) GetSingleDataModelFields(dbConn *models.DBConnection, schema string, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	// get fields
	query := fmt.Sprintf(`
		SELECT ordinal_position, column_name, data_type, is_nullable, column_default, character_maximum_length
		FROM information_schema.columns
		WHERE table_schema = '%s' AND table_name = '%s'
		ORDER BY ordinal_position;`,
		schema, name)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	fieldsData := data["rows"].([]map[string]interface{})
	// get constraints
	query = fmt.Sprintf(`SELECT conkey, conname, contype
		FROM pg_constraint WHERE conrelid = '"%s"."%s"'::regclass;`, schema, name)
	data, err = pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	constraintsData := data["rows"].([]map[string]interface{})
	return pgxutils.QueryToDataModel(fieldsData, constraintsData), err
}

func (pgqe *PostgresQueryEngine) GetSingleDataModelIndexes(dbConn *models.DBConnection, schema string, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`SELECT indexname, indexdef FROM pg_indexes
	WHERE schemaname = '%s' AND tablename = '%s';`, schema, name)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	returnedData := data["rows"].([]map[string]interface{})
	return returnedData, err
}

func (pgqe *PostgresQueryEngine) AddSingleDataModelColumn(dbConn *models.DBConnection, schema, name, columnName, dataType string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`ALTER TABLE %s.%s ADD COLUMN %s %s;`, schema, name, columnName, dataType)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) DeleteSingleDataModelColumn(dbConn *models.DBConnection, schema, name, columnName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`ALTER TABLE %s.%s DROP COLUMN %s;`, schema, name, columnName)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) GetData(dbConn *models.DBConnection, schema string, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string, config *models.QueryConfig) (map[string]interface{}, error) {
	sortQuery := ""
	if len(sort) == 2 {
		sortQuery = fmt.Sprintf(` ORDER BY %s %s`, sort[0], sort[1])
	}
	query := fmt.Sprintf(`SELECT ctid, * FROM "%s"."%s"%s LIMIT %d OFFSET %d;`, schema, name, sortQuery, limit, offset)
	countQuery := fmt.Sprintf(`SELECT count(*) FROM "%s"."%s";`, schema, name)
	if len(filter) > 1 {
		filter2 := ""
		if len(filter) == 3 {
			filter2 = " '" + filter[2] + "'"
		}
		query = fmt.Sprintf(`SELECT ctid, * FROM "%s"."%s" WHERE "%s" %s%s%s LIMIT %d OFFSET %d;`,
			schema,
			name,
			filter[0],
			filter[1],
			filter2,
			sortQuery,
			limit,
			offset)
		countQuery = fmt.Sprintf(`SELECT count(*) FROM "%s"."%s" WHERE "%s" %s%s;`,
			schema,
			name,
			filter[0],
			filter[1],
			filter2)
	}
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	if fetchCount {
		countData, err := pgqe.RunQuery(dbConn, countQuery, config)
		if err != nil {
			return nil, err
		}
		data["count"] = countData["rows"].([]map[string]interface{})[0]["0"]
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) UpdateSingleData(dbConn *models.DBConnection, schema string, name string, ctid string, columnName string, value string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`UPDATE "%s"."%s" SET "%s" = '%s' WHERE ctid = '%s' RETURNING ctid;`, schema, name, columnName, value, ctid)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	ctID := data["rows"].([]map[string]interface{})[0]["0"]
	data = map[string]interface{}{
		"ctid": ctID,
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) AddData(dbConn *models.DBConnection, schema string, name string, data map[string]interface{}, config *models.QueryConfig) (map[string]interface{}, error) {
	keys := []string{}
	values := []string{}
	for key, value := range data {
		keys = append(keys, key)
		val := value.(string)
		values = append(values, val)
	}
	keysStr := strings.Join(keys, ", ")
	valuesStr := strings.Join(values, "','")
	query := fmt.Sprintf(`INSERT INTO "%s"."%s"(%s) VALUES('%s') RETURNING ctid, *;`, schema, name, keysStr, valuesStr)
	rData, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	ctID := rData["rows"].([]map[string]interface{})[0]["0"]
	ndata := rData["rows"].([]map[string]interface{})[0]
	delete(ndata, "0")
	rData = map[string]interface{}{
		"ctid": ctID,
		"data": ndata,
	}
	return rData, err
}

func (pgqe *PostgresQueryEngine) DeleteData(dbConn *models.DBConnection, schema string, name string, ctids []string, config *models.QueryConfig) (map[string]interface{}, error) {
	ctidsStr := strings.Join(ctids, "', '")
	query := fmt.Sprintf(`DELETE FROM "%s"."%s" WHERE ctid IN ('%s');`, schema, name, ctidsStr)
	return pgqe.RunQuery(dbConn, query, config)
}

func (pgqe *PostgresQueryEngine) AddSingleDataModelIndex(dbConn *models.DBConnection, schema, name, indexName string, colNames []string, isUnique bool, config *models.QueryConfig) (map[string]interface{}, error) {
	isUniqueStr := ""
	if isUnique {
		isUniqueStr = "UNIQUE "
	}
	query := fmt.Sprintf(`CREATE %sINDEX %s ON "%s"."%s" (%s);`, isUniqueStr, indexName, schema, name, strings.Join(colNames, ", "))
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (pgqe *PostgresQueryEngine) DeleteSingleDataModelIndex(dbConn *models.DBConnection, indexName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`DROP INDEX %s;`, indexName)
	data, err := pgqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}
