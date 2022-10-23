package mongoqueryengine

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo/options"
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
	queryType := mongoutils.GetMongoQueryType(query)
	if queryType.QueryType == mongoutils.QUERY_FINDONE {
		result := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			FindOne(context.Background(), queryType.Data)
		if result.Err() != nil {
			return nil, result.Err()
		}
		keys, data := mongoutils.MongoSingleResultToJson(result)
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_FIND {
		cursor, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			Find(context.Background(), queryType.Data, &options.FindOptions{Limit: queryType.Limit, Skip: queryType.Skip})
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
	} else if queryType.QueryType == mongoutils.QUERY_RUNCMD {
		result := conn.Database(string(dbConn.DBName)).RunCommand(context.Background(), queryType.Data)
		if result.Err() != nil {
			return nil, result.Err()
		}
		keys, data := mongoutils.MongoSingleResultToJson(result)
		if createLog {
			queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
			go dbQueryLogDao.CreateDBQueryLog(queryLog)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_LISTCOLLECTIONS {
		list, err := conn.Database(string(dbConn.DBName)).ListCollectionNames(context.Background(), queryType.Data)
		if err != nil {
			return nil, err
		}
		if createLog {
			queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
			go dbQueryLogDao.CreateDBQueryLog(queryLog)
		}
		data := []map[string]interface{}{}
		for _, name := range list {
			data = append(data, map[string]interface{}{"collectionName": name})
		}
		return map[string]interface{}{
			"keys": []string{"collectionName"},
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_COUNT {
		count, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			CountDocuments(context.Background(), queryType.Data)
		if err != nil {
			return nil, err
		}
		if createLog {
			queryLog := models.NewQueryLog(user.ID, dbConn.ID, query)
			go dbQueryLogDao.CreateDBQueryLog(queryLog)
		}
		return map[string]interface{}{
			"keys": []string{"count"},
			"data": []map[string]interface{}{
				{
					"count": count,
				},
			},
		}, nil
	}
	return nil, errors.New("unknown query")
}

func (mqe *MongoQueryEngine) TestConnection(user *models.User, dbConn *models.DBConnection) bool {
	query := "db.runCommand({ping: 1})"
	data, err := mqe.RunQuery(user, dbConn, query, false)
	if err != nil {
		return false
	}
	test := data["data"].([]map[string]interface{})[0]["ok"].(float64)
	return test == 1
}

func (mqe *MongoQueryEngine) GetDataModels(user *models.User, dbConn *models.DBConnection) ([]map[string]interface{}, error) {
	query := "db.getCollectionNames()"
	data, err := mqe.RunQuery(user, dbConn, query, true)
	if err != nil {
		return nil, err
	}
	rdata := data["data"].([]map[string]interface{})
	return rdata, nil
}

func (mqe *MongoQueryEngine) GetData(user *models.User, dbConn *models.DBConnection, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string) (map[string]interface{}, error) {
	// sortQuery := ""
	// if len(sort) == 2 {
	// update sort query
	// }
	query := fmt.Sprintf(`db.%s.find().limit(%d).skip(%d)`, name, limit, offset)
	countQuery := fmt.Sprintf(`db.%s.count()`, name)
	// if len(filter) > 1 {
	// 	//update query & countQuery
	// }
	data, err := mqe.RunQuery(user, dbConn, query, false)
	if err != nil {
		return nil, err
	}
	if fetchCount {
		countData, err := mqe.RunQuery(user, dbConn, countQuery, false)
		if err != nil {
			return nil, err
		}
		data["count"] = countData["data"].([]map[string]interface{})[0]["count"]
	}
	return data, err
}
