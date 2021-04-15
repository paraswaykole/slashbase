package queryengines

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"slashbase.com/backend/config"
	"slashbase.com/backend/models"
)

type PostgresQueryEngine struct{}

func (pgqe PostgresQueryEngine) RunQuery(dbConn *models.DBConnection, query string) (map[string]interface{}, error) {
	port, err := strconv.Atoi(string(dbConn.DBPort))
	postBody, _ := json.Marshal(map[string]interface{}{
		"query":       query,
		"host":        dbConn.DBHost,
		"port":        port,
		"database":    dbConn.DBName,
		"user":        dbConn.DBUser,
		"password":    dbConn.DBPassword,
		"useSSH":      dbConn.UseSSH,
		"sshHost":     dbConn.SSHHost,
		"sshUser":     dbConn.SSHUser,
		"sshPassword": dbConn.SSHPassword,
		"sshKeyFile":  dbConn.SSHKeyFile,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(config.GetQueryEngineURLHost(), "application/json", responseBody)
	if err != nil {
		return nil, errors.New("An Error Occured")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	return data, err
}
