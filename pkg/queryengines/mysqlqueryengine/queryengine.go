package mysqlqueryengine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/mysqlqueryengine/mysqlutils"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
)

type MysqlQueryEngine struct {
	openConnections map[string]mysqlInstance
	mutex           *sync.Mutex
}

func InitMysqlQueryEngine() *MysqlQueryEngine {
	return &MysqlQueryEngine{
		openConnections: map[string]mysqlInstance{},
		mutex:           &sync.Mutex{},
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

func (mqe *MysqlQueryEngine) TestConnection(dbConn *models.DBConnection, config *models.QueryConfig) bool {
	query := "SELECT 1 AS test;"
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return false
	}
	test := data["rows"].([]map[string]interface{})[0]["0"].(int64)
	return test == 1
}

func (mqe *MysqlQueryEngine) GetDataModels(dbConn *models.DBConnection, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema='%s';", dbConn.DBName)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	rdata := data["rows"].([]map[string]interface{})
	return rdata, nil
}

func (mqe *MysqlQueryEngine) GetSingleDataModelFields(dbConn *models.DBConnection, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	// get fields
	query := fmt.Sprintf(`
		SELECT ordinal_position, column_name, data_type, is_nullable, column_default, character_maximum_length
		FROM information_schema.columns
		WHERE table_schema = '%s' AND table_name = '%s'
		ORDER BY ordinal_position;`,
		dbConn.DBName, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	fieldsData := data["rows"].([]map[string]interface{})
	// get constraints
	query = fmt.Sprintf(`SELECT * FROM information_schema.key_column_usage WHERE table_schema='%s' and table_name='%s'`, dbConn.DBName, name)
	data, err = mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	constraintsData := data["rows"].([]map[string]interface{})
	return mysqlutils.QueryToDataModel(fieldsData, constraintsData), err
}

func (mqe *MysqlQueryEngine) GetSingleDataModelIndexes(dbConn *models.DBConnection, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`SELECT statistics.index_name, GROUP_CONCAT(distinct column_name SEPARATOR ',') AS columns FROM information_schema.statistics WHERE table_schema = '%s' AND table_name = '%s' GROUP BY(statistics.index_name);`, dbConn.DBName, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	returnedData := data["rows"].([]map[string]interface{})
	return returnedData, err
}

func (mqe *MysqlQueryEngine) AddSingleDataModelColumn(dbConn *models.DBConnection, name, columnName, dataType string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`ALTER TABLE %s ADD COLUMN %s %s;`, name, columnName, dataType)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (mqe *MysqlQueryEngine) DeleteSingleDataModelColumn(dbConn *models.DBConnection, name, columnName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`ALTER TABLE %s DROP COLUMN %s;`, name, columnName)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (mqe *MysqlQueryEngine) GetData(dbConn *models.DBConnection, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string, config *models.QueryConfig) (map[string]interface{}, error) {
	sortQuery := ""
	if len(sort) == 2 {
		sortQuery = fmt.Sprintf(` ORDER BY %s %s`, sort[0], sort[1])
	}
	query := fmt.Sprintf(`SELECT * FROM %s%s LIMIT %d OFFSET %d;`, name, sortQuery, limit, offset)
	countQuery := fmt.Sprintf(`SELECT count(*) FROM %s;`, name)
	if len(filter) > 1 {
		filter2 := ""
		if len(filter) == 3 {
			filter2 = " '" + filter[2] + "'"
		}
		query = fmt.Sprintf(`SELECT * FROM %s WHERE %s %s%s%s LIMIT %d OFFSET %d;`,
			name,
			filter[0],
			filter[1],
			filter2,
			sortQuery,
			limit,
			offset)
		countQuery = fmt.Sprintf(`SELECT count(*) FROM %s WHERE %s %s%s;`,
			name,
			filter[0],
			filter[1],
			filter2)
	}
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	if fetchCount {
		countData, err := mqe.RunQuery(dbConn, countQuery, config)
		if err != nil {
			return nil, err
		}
		data["count"] = countData["rows"].([]map[string]interface{})[0]["0"]
	}
	return data, err
}

func (mqe *MysqlQueryEngine) UpdateSingleData(dbConn *models.DBConnection, name string, pkey string, columnName string, value string, config *models.QueryConfig) (map[string]interface{}, error) {
	return nil, errors.New("not implemented")
}

func (mqe *MysqlQueryEngine) AddData(dbConn *models.DBConnection, name string, data map[string]interface{}, config *models.QueryConfig) (map[string]interface{}, error) {
	keys := []string{}
	values := []string{}
	for key, value := range data {
		keys = append(keys, key)
		val := value.(string)
		values = append(values, val)
	}
	keysStr := strings.Join(keys, ", ")
	valuesStr := strings.Join(values, "','")
	query := fmt.Sprintf(`INSERT INTO %s(%s) VALUES('%s');`, name, keysStr, valuesStr)
	resultData, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"data":    data,
		"message": resultData["message"].(string),
	}, err
}

func (mqe *MysqlQueryEngine) DeleteData(dbConn *models.DBConnection, name string, ctids []string, config *models.QueryConfig) (map[string]interface{}, error) {
	return nil, errors.New("not implemented")
}

func (mqe *MysqlQueryEngine) AddSingleDataModelIndex(dbConn *models.DBConnection, name, indexName string, colNames []string, isUnique bool, config *models.QueryConfig) (map[string]interface{}, error) {
	isUniqueStr := ""
	if isUnique {
		isUniqueStr = "UNIQUE "
	}
	query := fmt.Sprintf(`CREATE %sINDEX %s ON %s (%s);`, isUniqueStr, indexName, name, strings.Join(colNames, ", "))
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (mqe *MysqlQueryEngine) DeleteSingleDataModelIndex(dbConn *models.DBConnection, name, indexName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf("DROP INDEX `%s` ON `%s`;", indexName, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}
