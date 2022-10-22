package mongoqueryengine

import (
	"context"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
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
	queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
	queryType := mongoutils.GetMongoQueryType(query)
	if queryType.QueryType == mongoutils.QUERY_FINDONE {
		result := conn.Database(string(dbConn.DBName)).Collection(queryType.CollectionName).FindOne(context.Background(), queryType.Filter)
		if result.Err() != nil {
			return nil, result.Err()
		}
		var bsonData bson.D
		err = result.Decode(&bsonData)
		if err != nil {
			return nil, err
		}
		data = bsonData.Map()
	} else if queryType.QueryType == mongoutils.QUERY_FIND {
		_, err := conn.Database(string(dbConn.DBName)).Collection(queryType.CollectionName).Find(context.Background(), queryType.Filter)
		if err != nil {
			return nil, err
		}
		// TODO: read cursor
	}
	if createLog {
		go dbQueryLogDao.CreateDBQueryLog(queryLog)
	}

	return data, nil
}
