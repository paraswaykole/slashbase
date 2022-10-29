package mongoqueryengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
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
			FindOne(context.Background(), queryType.Args[0])
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
			Find(context.Background(), queryType.Args[0], &options.FindOptions{Limit: queryType.Limit, Skip: queryType.Skip})
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
	} else if queryType.QueryType == mongoutils.QUERY_INSERTONE {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			InsertOne(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"insertedId"},
			"data": []map[string]interface{}{
				{
					"insertedId": result.InsertedID,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_INSERT {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			InsertMany(context.Background(), queryType.Args[0].(bson.A))
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"insertedIDs"},
			"data": []map[string]interface{}{
				{
					"insertedIDs": result.InsertedIDs,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_DELETEONE {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			DeleteOne(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"deletedCount"},
			"data": []map[string]interface{}{
				{
					"deletedCount": result.DeletedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_DELETEMANY {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			DeleteMany(context.Background(), queryType.Args[0].(bson.D))
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"deletedCount"},
			"data": []map[string]interface{}{
				{
					"deletedCount": result.DeletedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_UPDATEONE {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			UpdateOne(context.Background(), queryType.Args[0], queryType.Args[1])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"updatedCount", "upsertedCount"},
			"data": []map[string]interface{}{
				{
					"updatedCount":  result.ModifiedCount,
					"upsertedCount": result.UpsertedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_UPDATEMANY {
		result, err := conn.Database(string(dbConn.DBName)).
			Collection(queryType.CollectionName).
			UpdateMany(context.Background(), queryType.Args[0], queryType.Args[1])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"keys": []string{"updatedCount", "upsertedCount"},
			"data": []map[string]interface{}{
				{
					"updatedCount":  result.ModifiedCount,
					"upsertedCount": result.UpsertedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_RUNCMD {
		result := conn.Database(string(dbConn.DBName)).RunCommand(context.Background(), queryType.Args[0])
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
		list, err := conn.Database(string(dbConn.DBName)).ListCollectionNames(context.Background(), queryType.Args[0])
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
			CountDocuments(context.Background(), queryType.Args[0])
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
	data, err := mqe.RunQuery(user, dbConn, query, true)
	if err != nil {
		return nil, err
	}
	if fetchCount {
		countData, err := mqe.RunQuery(user, dbConn, countQuery, true)
		if err != nil {
			return nil, err
		}
		data["count"] = countData["data"].([]map[string]interface{})[0]["count"]
	}
	return data, err
}

func (mqe *MongoQueryEngine) UpdateSingleData(user *models.User, dbConn *models.DBConnection, name string, underscoreID string, documentData string) (map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.updateOne({_id: ObjectId("%s")}, {$set: %s } )`, name, underscoreID, documentData)
	data, err := mqe.RunQuery(user, dbConn, query, true)
	if err != nil {
		return nil, err
	}
	updatedCount := data["data"].([]map[string]interface{})[0]["updatedCount"]
	data = map[string]interface{}{
		"updatedCount": updatedCount,
	}
	return data, err
}

func (mqe *MongoQueryEngine) AddData(user *models.User, dbConn *models.DBConnection, schema string, name string, data map[string]interface{}) (map[string]interface{}, error) {
	dataStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`db.%s.insertOne(%s)`, name, string(dataStr))
	rData, err := mqe.RunQuery(user, dbConn, query, true)
	if err != nil {
		return nil, err
	}
	insertedID := rData["data"].([]map[string]interface{})[0]["insertedId"]
	rData = map[string]interface{}{
		"insertedId": insertedID,
	}
	return rData, err
}

func (mqe *MongoQueryEngine) DeleteData(user *models.User, dbConn *models.DBConnection, name string, underscoreIds []string) (map[string]interface{}, error) {
	for i, id := range underscoreIds {
		underscoreIds[i] = fmt.Sprintf(`ObjectId("%s")`, id)
	}
	underscoreIdsStr := strings.Join(underscoreIds, ", ")
	query := fmt.Sprintf(`db.%s.deleteMany({ _id : { "$in" : [%s] }})`, name, underscoreIdsStr)
	fmt.Println(query)
	return mqe.RunQuery(user, dbConn, query, true)
}
