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
	query := fmt.Sprintf(`SELECT DISTINCT table_name, index_name FROM information_schema.statistics WHERE table_schema = '%s' AND table_name = '%s';`, dbConn.DBName, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	returnedData := data["rows"].([]map[string]interface{})
	return returnedData, err
}
