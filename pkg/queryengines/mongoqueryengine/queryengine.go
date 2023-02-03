package mongoqueryengine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/slashbaseide/slashbase/pkg/queryengines/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/mongoqueryengine/mongoutils"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoQueryEngine struct {
	openClients map[string]mongoClientInstance
	mutex       *sync.Mutex
}

func InitMongoQueryEngine() *MongoQueryEngine {
	return &MongoQueryEngine{
		openClients: map[string]mongoClientInstance{},
		mutex:       &sync.Mutex{},
	}
}

func (mqe *MongoQueryEngine) RunQuery(dbConn *models.DBConnection, query string, config *models.QueryConfig) (map[string]interface{}, error) {
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
	conn, err := mqe.getConnection(dbConn.ID, string(dbConn.DBScheme), string(dbConn.DBHost), uint16(port), string(dbConn.DBUser), string(dbConn.DBPassword), dbConn.UseSSL)
	if err != nil {
		return nil, err
	}
	db := conn.Database(string(dbConn.DBName))
	queryType := mongoutils.GetMongoQueryType(query)

	queryTypeRead := mongoutils.IsQueryTypeRead(queryType)
	if !queryTypeRead && config.ReadOnly {
		return nil, errors.New("not allowed run this query")
	}

	if queryType.QueryType == mongoutils.QUERY_FINDONE {
		result := db.Collection(queryType.CollectionName).
			FindOne(context.Background(), queryType.Args[0])
		if result.Err() != nil {
			return nil, result.Err()
		}
		keys, data := mongoutils.MongoSingleResultToJson(result)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_FIND {
		collection := db.Collection(queryType.CollectionName)
		opts := &options.FindOptions{Limit: queryType.Limit, Skip: queryType.Skip, Sort: queryType.Sort}
		if len(queryType.Args) > 1 {
			opts.SetProjection(queryType.Args[1])
		}
		cursor, err := collection.
			Find(context.Background(), queryType.Args[0], opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		keys, data := mongoutils.MongoCursorToJson(cursor)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_INSERTONE {
		result, err := db.Collection(queryType.CollectionName).
			InsertOne(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
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
		result, err := db.Collection(queryType.CollectionName).
			InsertMany(context.Background(), queryType.Args[0].(bson.A))
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
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
		result, err := db.Collection(queryType.CollectionName).
			DeleteOne(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
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
		result, err := db.Collection(queryType.CollectionName).
			DeleteMany(context.Background(), queryType.Args[0].(bson.D))
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
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
		result, err := db.Collection(queryType.CollectionName).
			UpdateOne(context.Background(), queryType.Args[0], queryType.Args[1])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": []string{"matchedCount", "updatedCount", "upsertedCount"},
			"data": []map[string]interface{}{
				{
					"matchedCount":  result.MatchedCount,
					"updatedCount":  result.ModifiedCount,
					"upsertedCount": result.UpsertedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_UPDATEMANY {
		result, err := db.Collection(queryType.CollectionName).
			UpdateMany(context.Background(), queryType.Args[0], queryType.Args[1])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": []string{"matchedCount", "updatedCount", "upsertedCount"},
			"data": []map[string]interface{}{
				{
					"matchedCount":  result.MatchedCount,
					"updatedCount":  result.ModifiedCount,
					"upsertedCount": result.UpsertedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_REPLACEONE {
		result, err := db.Collection(queryType.CollectionName).
			ReplaceOne(context.Background(), queryType.Args[0], queryType.Args[1])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": []string{"matchedCount", "updatedCount", "upsertedCount"},
			"data": []map[string]interface{}{
				{
					"matchedCount":  result.MatchedCount,
					"updatedCount":  result.ModifiedCount,
					"upsertedCount": result.UpsertedCount,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_RUNCMD {
		result := db.RunCommand(context.Background(), queryType.Args[0])
		if result.Err() != nil {
			return nil, result.Err()
		}
		keys, data := mongoutils.MongoSingleResultToJson(result)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_GETINDEXES {
		cursor, err := db.RunCommandCursor(context.Background(), bson.M{
			"listIndexes": queryType.CollectionName,
		})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		keys, data := mongoutils.MongoCursorToJson(cursor)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_LISTCOLLECTIONS {
		list, err := db.ListCollectionNames(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
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
		opts := &options.CountOptions{Limit: queryType.Limit, Skip: queryType.Skip}
		count, err := db.Collection(queryType.CollectionName).
			CountDocuments(context.Background(), queryType.Args[0], opts)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": []string{"count"},
			"data": []map[string]interface{}{
				{
					"count": count,
				},
			},
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_AGGREGATE {
		cursor, err := db.Collection(queryType.CollectionName).
			Aggregate(context.Background(), queryType.Args[0])
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		keys, data := mongoutils.MongoCursorToJson(cursor)
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"keys": keys,
			"data": data,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_DROP {
		err := db.Collection(queryType.CollectionName).Drop(context.Background())
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"message": "droped collection: " + queryType.CollectionName,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_CREATEINDEX {
		opts := options.Index()
		if len(queryType.Args) > 1 {
			if d, ok := queryType.Args[1].(bson.D); ok {
				if unique, ok := d.Map()["unique"].(bool); ok {
					opts.SetUnique(unique)
				}
				if name, ok := d.Map()["name"].(string); ok {
					opts.SetName(name)
				}
			}
		}
		indexModel := mongo.IndexModel{
			Keys:    queryType.Args[0],
			Options: opts,
		}
		idxName, err := db.Collection(queryType.CollectionName).Indexes().CreateOne(context.Background(), indexModel)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"message": "created index: " + idxName,
		}, nil
	} else if queryType.QueryType == mongoutils.QUERY_DROPINDEX {
		indexName, ok := queryType.Args[0].(string)
		if !ok {
			return nil, errors.New("invalid query")
		}
		if indexName == "*" {
			_, err := db.Collection(queryType.CollectionName).Indexes().DropAll(context.Background(), options.DropIndexes())
			if err != nil {
				return nil, err
			}
			if config.CreateLogFn != nil {
				config.CreateLogFn(query)
			}
			return map[string]interface{}{
				"message": "droped all indexes in collection: " + queryType.CollectionName,
			}, nil
		}
		_, err := db.Collection(queryType.CollectionName).Indexes().DropOne(context.Background(), indexName, options.DropIndexes())
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return map[string]interface{}{
			"message": "droped index: " + indexName,
		}, nil
	}
	return nil, errors.New("unknown query")
}

func (mqe *MongoQueryEngine) TestConnection(dbConn *models.DBConnection, config *models.QueryConfig) bool {
	query := "db.runCommand({ping: 1})"
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return false
	}
	test := data["data"].([]map[string]interface{})[0]["ok"].(float64)
	return test == 1
}

func (mqe *MongoQueryEngine) GetDataModels(dbConn *models.DBConnection, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := "db.getCollectionNames()"
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	rdata := data["data"].([]map[string]interface{})
	return rdata, nil
}

func (mqe *MongoQueryEngine) GetSingleDataModelFields(dbConn *models.DBConnection, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.aggregate([{$sample: {size: 1000}}])`, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	returnedKeys := data["keys"].([]string)
	returnedData := data["data"].([]map[string]interface{})
	return mongoutils.AnalyseFieldsSchema(returnedKeys, returnedData), err
}

func (mqe *MongoQueryEngine) GetSingleDataModelIndexes(dbConn *models.DBConnection, name string, config *models.QueryConfig) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.getIndexes()`, name)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	returnedData := data["data"].([]map[string]interface{})
	return mongoutils.GetCollectionIndexes(returnedData), err
}

func (mqe *MongoQueryEngine) AddSingleDataModelKey(dbConn *models.DBConnection, schema, name, columnName, dataType string) (map[string]interface{}, error) {
	return nil, errors.New("not supported yet")
}

func (mqe *MongoQueryEngine) DeleteSingleDataModelKey(dbConn *models.DBConnection, schema, name, columnName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.updateMany({}, {$unset: {%s: ""}})`, name, columnName)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (mqe *MongoQueryEngine) GetData(dbConn *models.DBConnection, name string, limit int, offset int64, fetchCount bool, filter []string, sort []string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.find().limit(%d).skip(%d)`, name, limit, offset)
	countQuery := fmt.Sprintf(`db.%s.count()`, name)
	if len(filter) == 1 && strings.HasPrefix(filter[0], "{") && strings.HasSuffix(filter[0], "}") {
		query = fmt.Sprintf(`db.%s.find(%s).limit(%d).skip(%d)`, name, filter[0], limit, offset)
		countQuery = fmt.Sprintf(`db.%s.count(%s)`, name, filter[0])
	}
	if len(sort) == 1 {
		query = query + fmt.Sprintf(`.sort(%s)`, sort[0])
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
		data["count"] = countData["data"].([]map[string]interface{})[0]["count"]
	}
	return data, err
}

func (mqe *MongoQueryEngine) UpdateSingleData(dbConn *models.DBConnection, name string, underscoreID string, documentData string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.replaceOne({_id: ObjectId("%s")}, %s )`, name, underscoreID, documentData)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	updatedCount := data["data"].([]map[string]interface{})[0]["updatedCount"]
	data = map[string]interface{}{
		"updatedCount": updatedCount,
	}
	return data, err
}

func (mqe *MongoQueryEngine) AddData(dbConn *models.DBConnection, schema string, name string, data map[string]interface{}, config *models.QueryConfig) (map[string]interface{}, error) {
	dataStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`db.%s.insertOne(%s)`, name, string(dataStr))
	rData, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	insertedID := rData["data"].([]map[string]interface{})[0]["insertedId"]
	rData = map[string]interface{}{
		"insertedId": insertedID,
	}
	return rData, err
}

func (mqe *MongoQueryEngine) DeleteData(dbConn *models.DBConnection, name string, underscoreIds []string, config *models.QueryConfig) (map[string]interface{}, error) {
	for i, id := range underscoreIds {
		underscoreIds[i] = fmt.Sprintf(`ObjectId("%s")`, id)
	}
	underscoreIdsStr := strings.Join(underscoreIds, ", ")
	query := fmt.Sprintf(`db.%s.deleteMany({ _id : { "$in" : [%s] }})`, name, underscoreIdsStr)
	fmt.Println(query)
	return mqe.RunQuery(dbConn, query, config)
}

func (mqe *MongoQueryEngine) AddSingleDataModelIndex(dbConn *models.DBConnection, name, indexName string, keyNames []string, isUnique bool, config *models.QueryConfig) (map[string]interface{}, error) {
	for i := range keyNames {
		keyNames[i] = keyNames[i] + ": 1"
	}
	query := fmt.Sprintf(`db.%s.createIndex({%s}, {unqiue: %t, name: %s});`, name, strings.Join(keyNames, ", "), isUnique, indexName)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (mqe *MongoQueryEngine) DeleteSingleDataModelIndex(dbConn *models.DBConnection, name, indexName string, config *models.QueryConfig) (map[string]interface{}, error) {
	query := fmt.Sprintf(`db.%s.dropIndex("%s")`, name, indexName)
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return nil, err
	}
	return data, err
}
