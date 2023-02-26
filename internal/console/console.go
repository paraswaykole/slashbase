package console

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

func HandleCommand(dbConnection *models.DBConnection, cmdString string, queryConfigs *qemodels.QueryConfig) string {

	if cmdString == "help" {
		return helpText(dbConnection)
	}

	if cmdString == "ping" {
		err := queryengines.TestConnection(dbConnection.ToQEConnection(), queryConfigs)
		if err == nil {
			return "pong"
		}
		return "unable to connect to db"
	}

	result, err := queryengines.RunQuery(dbConnection.ToQEConnection(), cmdString, queryConfigs)
	if err != nil {
		return fmt.Sprintf("error: '%s'\n", err.Error())
	}
	if dbConnection.Type == qemodels.DBTYPE_POSTGRES {
		return postgresResult(result)
	} else if dbConnection.Type == qemodels.DBTYPE_MYSQL {
		return mysqlResult(result)
	} else if dbConnection.Type == qemodels.DBTYPE_MONGO {
		return mongoResult(result)
	}

	return "unknown command"
}

func helpText(dbConnection *models.DBConnection) string {
	response := "Type 'ping' to test connection. If it returns pong, connection is successful."

	if dbConnection.Type == qemodels.DBTYPE_POSTGRES {
		response += "\nOr type postgresql command and press enter to run it."
		response += "\nFor example: 'SELECT * FROM <table name>;'"
	} else if dbConnection.Type == qemodels.DBTYPE_MYSQL {
		response += "\nOr type mysql command and press enter to run it."
		response += "\nFor example: 'SELECT * FROM <table name>;'"
	} else if dbConnection.Type == qemodels.DBTYPE_MONGO {
		response += "\nOr type mongo command and press enter to run it."
		response += "\nFor example: 'db.<collection name>.find({});'"
	}

	return response
}

func postgresResult(data map[string]interface{}) string {

	if msg, ok := data["message"].(string); ok {
		return fmt.Sprintf("Result: '%s'\n", msg)
	}

	t := table.NewWriter()

	headers := table.Row{}
	for _, colName := range data["columns"].([]string) {
		headers = append(headers, colName)
	}

	allRows := []table.Row{}
	for _, rdata := range data["rows"].([]map[string]interface{}) {
		row := make(table.Row, len(rdata))
		for key, value := range rdata {
			idx, _ := strconv.Atoi(key)
			if value == nil {
				row[idx] = "NULL"
			} else {
				row[idx] = value
			}
		}
		allRows = append(allRows, row)
	}

	defStyle := table.StyleDefault
	defStyle.Format.Header = text.FormatDefault
	t.SetStyle(defStyle)
	t.AppendHeader(headers)
	t.AppendRows(allRows)
	return t.Render()
}

func mysqlResult(data map[string]interface{}) string {

	if msg, ok := data["message"].(string); ok {
		return fmt.Sprintf("Result: '%s'\n", msg)
	}

	t := table.NewWriter()

	headers := table.Row{}
	for _, colName := range data["columns"].([]string) {
		headers = append(headers, colName)
	}

	allRows := []table.Row{}
	for _, rdata := range data["rows"].([]map[string]interface{}) {
		row := make(table.Row, len(rdata))
		for key, value := range rdata {
			idx, _ := strconv.Atoi(key)
			if value == nil {
				row[idx] = "NULL"
			} else {
				row[idx] = value
			}
		}
		allRows = append(allRows, row)
	}

	defStyle := table.StyleDefault
	defStyle.Format.Header = text.FormatDefault
	t.SetStyle(defStyle)
	t.AppendHeader(headers)
	t.AppendRows(allRows)
	return t.Render()
}

func mongoResult(data map[string]interface{}) string {

	if msg, ok := data["message"].(string); ok {
		return fmt.Sprintf("Result: '%s'\n", msg)
	}

	allRows := []table.Row{}
	for _, rdata := range data["data"].([]map[string]interface{}) {
		row, _ := json.MarshalIndent(rdata, "", " ")
		allRows = append(allRows, table.Row{string(row)})
	}

	t := table.NewWriter()
	defStyle := table.StyleDefault
	defStyle.Options.SeparateRows = true
	t.SetStyle(defStyle)
	t.SetOutputMirror(os.Stdout)
	t.AppendRows(allRows)
	return t.Render()
}
