package mongoqueryengine

import (
	"context"
	"fmt"
	"strconv"

	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/internal/models/sbsql"
	"slashbase.com/backend/pkg/queryengines/mongoqueryengine/mongoutils"
	"slashbase.com/backend/pkg/sshtunnel"
)

var dbQueryLogDao daos.DBQueryLogDao

type MongoQueryEngine struct {
	openClients map[string]mongoClientInstance
}

func InitMongoQueryEngine() *MongoQueryEngine {
	return &MongoQueryEngine{
		openClients: map[string]mongoClientInstance{},
	}
}

func (mqe *MongoQueryEngine) RunQuery(user *models.User, dbConn *models.DBConnection, query string, createLog bool) (map[string]interface{}, error) {
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
		dbConn.DBHost = sbsql.CryptedData("localhost")
		dbConn.DBPort = sbsql.CryptedData(fmt.Sprintf("%d", sshTun.GetLocalEndpoint().Port))
	}
	port, _ = strconv.Atoi(string(dbConn.DBPort))
	conn, err := mqe.getConnection(dbConn.ConnectionUser.ID, string(dbConn.DBHost), uint16(port), string(dbConn.DBName), string(dbConn.ConnectionUser.DBUser), string(dbConn.ConnectionUser.DBPassword))
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	queryType := mongoutils.GetMongoQueryType(query)
	if queryType.QueryType == mongoutils.QUERY_FINDONE {
		result := conn.Database(string(dbConn.DBName)).Collection(queryType.CollectionName).FindOne(context.Background(), queryType.Filter)
		if result.Err() != nil {
			return nil, result.Err()
		}
		keys, data := mongoutils.MongoSingleResultToJson(result)
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_FIND {
		cursor, err := conn.Database(string(dbConn.DBName)).Collection(queryType.CollectionName).Find(context.Background(), queryType.Filter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		keys, data := mongoutils.MongoCursorToJson(cursor)
		if createLog {
			queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
			go dbQueryLogDao.CreateDBQueryLog(queryLog)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	}
	if createLog {
		queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
		go dbQueryLogDao.CreateDBQueryLog(queryLog)
	}

	return data, nil
}
