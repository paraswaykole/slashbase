package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gohxs/readline"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

var cliApp struct {
	CurrentDB *qemodels.DBConnection
}

func startCLI() {
	term, err := readline.NewEx(&readline.Config{
		Prompt: "slashbase > ",
		Output: display,
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		line, err := term.Readline()
		if err != nil {
			log.Fatal(err)
		}
		handleCmd(line)
		if cliApp.CurrentDB == nil {
			term.SetPrompt("slashbase > ")
		} else {
			term.SetPrompt(fmt.Sprintf("%s > ", cliApp.CurrentDB.Name))
		}
	}
}

func handleCmd(cmdText string) {

	if cmdText == "" {
		return
	}
	if cmdText == "exit" {
		os.Exit(1)
		return
	}

	if cmdText == "help" {
		printHelp()
		return
	}

	if strings.HasPrefix(cmdText, "\\base") {
		switchDB(cmdText)
	} else {
		runQuery(cmdText)
	}

}

func printHelp() {
	fmt.Println("To add a new database use the IDE interface running at https://localhost:" + config.GetConfig().Port)
	fmt.Println("To connect to existing db type '\\base db-nick-name'.")
	fmt.Println("Once connected to db, type your query and press enter to get query results.")
	fmt.Println("To end the program, type 'exit'.")
}

func switchDB(cmdText string) {
	dbname := strings.Replace(cmdText, "\\base ", "", 1)

	dbConn, err := dao.DBConnection.GetDBConnectionByName(dbname)
	if err != nil {
		fmt.Println("err: ", err.Error())
		fmt.Printf("no db found by name: '%s'\n", dbname)
		return
	}
	qeDBConn := dbConn.ToQEConnection()
	success := queryengines.TestConnection(qeDBConn, getQueryConfigs(qeDBConn))
	if !success {
		fmt.Printf("cannot connect to db: '%s'\n", dbname)
		return
	}

	cliApp.CurrentDB = qeDBConn
	fmt.Printf("connected to: '%s'\n", dbname)
}

func runQuery(queryCmd string) {
	if cliApp.CurrentDB == nil {
		fmt.Printf("not connected to any database\n")
		return
	}
	result, err := queryengines.RunQuery(cliApp.CurrentDB, queryCmd, getQueryConfigs(cliApp.CurrentDB))
	if err != nil {
		fmt.Printf("error: '%s'\n", err.Error())
		return
	}
	if cliApp.CurrentDB.Type == qemodels.DBTYPE_POSTGRES {
		postgresResult(result)
	} else if cliApp.CurrentDB.Type == qemodels.DBTYPE_MYSQL {
		mysqlResult(result)
	} else {
		mongoResult(result)
	}
}

func getQueryConfigs(dbConn *qemodels.DBConnection) *qemodels.QueryConfig {
	createLog := func(query string) {
		queryLog := models.NewQueryLog(dbConn.ID, query)
		go dao.DBQueryLog.CreateDBQueryLog(queryLog)
	}
	readOnly := false
	return qemodels.NewQueryConfig(readOnly, createLog)
}

func postgresResult(data map[string]interface{}) {

	if msg, ok := data["message"].(string); ok {
		fmt.Printf("Result: '%s'\n", msg)
		return
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
			row[idx] = value
		}
		allRows = append(allRows, row)
	}

	t.SetOutputMirror(os.Stdout)
	defStyle := table.StyleDefault
	defStyle.Format.Header = text.FormatDefault
	t.SetStyle(defStyle)
	t.AppendHeader(headers)
	t.AppendRows(allRows)
	t.Render()
}

func mysqlResult(data map[string]interface{}) {

	if msg, ok := data["message"].(string); ok {
		fmt.Printf("Result: '%s'\n", msg)
		return
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
			row[idx] = value
		}
		allRows = append(allRows, row)
	}

	t.SetOutputMirror(os.Stdout)
	defStyle := table.StyleDefault
	defStyle.Format.Header = text.FormatDefault
	t.SetStyle(defStyle)
	t.AppendHeader(headers)
	t.AppendRows(allRows)
	t.Render()
}

func mongoResult(data map[string]interface{}) {

	if msg, ok := data["message"].(string); ok {
		fmt.Printf("Result: '%s'\n", msg)
		return
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
	t.Render()

}
