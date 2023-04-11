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

func (mqe *MongoQueryEngine) runFindOneQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result := db.Collection(queryType.CollectionName).FindOne(context.Background(), queryType.Args[0])
	if result.Err() != nil {
		return nil, result.Err()
	}
	keys, data := mongoutils.MongoSingleResultToJson(result)
	return map[string]interface{}{
		"keys": keys,
		"data": data,
	}, nil
}

func (mqe *MongoQueryEngine) runFindQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
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
	return map[string]interface{}{
		"keys": keys,
		"data": data,
	}, nil

}

func (mqe *MongoQueryEngine) runInsertOneQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
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

}

func (mqe *MongoQueryEngine) runInsertManyQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
		InsertMany(context.Background(), queryType.Args[0].(bson.A))
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"keys": []string{"insertedIds"},
		"data": []map[string]interface{}{
			{
				"insertedIds": result.InsertedIDs,
			},
		},
	}, nil

}

func (mqe *MongoQueryEngine) runDeleteOneQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
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

}

func (mqe *MongoQueryEngine) runDeleteManyQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
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

}

func (mqe *MongoQueryEngine) runUpdateOneQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
		UpdateOne(context.Background(), queryType.Args[0], queryType.Args[1])
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"keys": []string{"matchedCount", "modifiedCount", "upsertedCount"},
		"data": []map[string]interface{}{
			{
				"matchedCount":  result.MatchedCount,
				"modifiedCount": result.ModifiedCount,
				"upsertedCount": result.UpsertedCount,
			},
		},
	}, nil

}

func (mwq *MongoQueryEngine) runUpdateManyQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
		UpdateMany(context.Background(), queryType.Args[0], queryType.Args[1])
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"keys": []string{"matchedCount", "modifiedCount", "upsertedCount"},
		"data": []map[string]interface{}{
			{
				"matchedCount":  result.MatchedCount,
				"modifiedCount": result.ModifiedCount,
				"upsertedCount": result.UpsertedCount,
			},
		},
	}, nil

}

func (mqe *MongoQueryEngine) runReplaceOneQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result, err := db.Collection(queryType.CollectionName).
		ReplaceOne(context.Background(), queryType.Args[0], queryType.Args[1])
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"keys": []string{"matchedCount", "modifiedCount", "upsertedCount"},
		"data": []map[string]interface{}{
			{
				"matchedCount":  result.MatchedCount,
				"modifiedCount": result.ModifiedCount,
				"upsertedCount": result.UpsertedCount,
			},
		},
	}, nil

}

func (mqe *MongoQueryEngine) runCMDQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	result := db.RunCommand(context.Background(), queryType.Args[0])
	if result.Err() != nil {
		return nil, result.Err()
	}
	keys, data := mongoutils.MongoSingleResultToJson(result)

	return map[string]interface{}{
			"keys": keys,
			"data": data,
		},
		nil

}

func (mqe *MongoQueryEngine) runGetIndexesQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	cursor, err := db.RunCommandCursor(context.Background(), bson.M{
		"listIndexes": queryType.CollectionName,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	keys, data := mongoutils.MongoCursorToJson(cursor)

	return map[string]interface{}{
		"keys": keys,
		"data": data,
	}, nil

}

func (mqe *MongoQueryEngine) runListCollectionsQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	list, err := db.ListCollectionNames(context.Background(), queryType.Args[0])
	if err != nil {
		return nil, err
	}
	data := []map[string]interface{}{}
	for _, name := range list {
		data = append(data, map[string]interface{}{"collectionName": name})
	}
	return map[string]interface{}{
		"keys": []string{"collectionName"},
		"data": data,
	}, nil

}

func (mqe *MongoQueryEngine) runCountQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	opts := &options.CountOptions{Limit: queryType.Limit, Skip: queryType.Skip}
	count, err := db.Collection(queryType.CollectionName).
		CountDocuments(context.Background(), queryType.Args[0], opts)
	if err != nil {
		return nil, err
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

func (mqe *MongoQueryEngine) runAggregateQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	cursor, err := db.Collection(queryType.CollectionName).
		Aggregate(context.Background(), queryType.Args[0])
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	keys, data := mongoutils.MongoCursorToJson(cursor)

	return map[string]interface{}{
		"keys": keys,
		"data": data,
	}, nil
}

func (mqe *MongoQueryEngine) runDropQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	err := db.Collection(queryType.CollectionName).Drop(context.Background())
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"message": "droped collection: " + queryType.CollectionName,
	}, nil
}

func (mqe *MongoQueryEngine) runCreateIndexQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
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

	return map[string]interface{}{
		"message": "created index: " + idxName,
	}, nil
}

func (mqe *MongoQueryEngine) runDropIndexQuery(db *mongo.Database, queryType *mongoutils.MongoQuery) (map[string]interface{}, error) {
	indexName, ok := queryType.Args[0].(string)
	if !ok {
		return nil, errors.New("invalid query")
	}
	if indexName == "*" {
		_, err := db.Collection(queryType.CollectionName).Indexes().DropAll(context.Background(), options.DropIndexes())
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"message": "droped all indexes in collection: " + queryType.CollectionName,
		}, nil
	}
	_, err := db.Collection(queryType.CollectionName).Indexes().DropOne(context.Background(), indexName, options.DropIndexes())
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "droped index: " + indexName,
	}, nil
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

	switch queryType.QueryType {

	case mongoutils.QUERY_FINDONE:
		result, err := mqe.runFindOneQuery(db, queryType)
		if err != nil {
			return nil, err
		}

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return result, nil
	case mongoutils.QUERY_FIND:
		result, err := mqe.runFindQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil

	case mongoutils.QUERY_INSERTONE:
		result, err := mqe.runInsertOneQuery(db, queryType)
		if err != nil {
			return nil, err
		}

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil
	case mongoutils.QUERY_INSERT:
		result, err := mqe.runInsertManyQuery(db, queryType)

		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil

	case mongoutils.QUERY_DELETEONE:
		result, err := mqe.runDeleteOneQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil

	case mongoutils.QUERY_DELETEMANY:
		result, err := mqe.runDeleteManyQuery(db, queryType)
		if err != nil {
			return nil, err
		}

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil
	case mongoutils.QUERY_UPDATEONE:
		result, err := mqe.runUpdateOneQuery(db, queryType)

		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return result, nil

	case mongoutils.QUERY_UPDATEMANY:
		result, err := mqe.runUpdateManyQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil

	case mongoutils.QUERY_REPLACEONE:
		result, err := mqe.runReplaceOneQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil

	case mongoutils.QUERY_RUNCMD:
		result, err := mqe.runCMDQuery(db, queryType)
		if err != nil {
			return nil, err
		}

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return result, nil
	case mongoutils.QUERY_GETINDEXES:
		results, err := mqe.runGetIndexesQuery(db, queryType)
		if err != nil {
			return nil, err
		}

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return results, nil
	case mongoutils.QUERY_LISTCOLLECTIONS:
		list, err := mqe.runListCollectionsQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return list, nil
	case mongoutils.QUERY_COUNT:

		count, err := mqe.runCountQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return count, nil
	case mongoutils.QUERY_AGGREGATE:
		cursor, err := mqe.runAggregateQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return cursor, nil

	case mongoutils.QUERY_DROP:
		result, err := mqe.runDropQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return result, nil
	case mongoutils.QUERY_CREATEINDEX:
		index, err := mqe.runCreateIndexQuery(db, queryType)
		if err != nil {
			return nil, err
		}
		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}

		return index, nil

	case mongoutils.QUERY_DROPINDEX:
		result, err := mqe.runDropIndexQuery(db, queryType)

		if config.CreateLogFn != nil {
			config.CreateLogFn(query)
		}
		return result, err

	default:
		return nil, errors.New("unknown query")
	}

}

func (mqe *MongoQueryEngine) TestConnection(dbConn *models.DBConnection, config *models.QueryConfig) error {
	query := "db.runCommand({ping: 1})"
	data, err := mqe.RunQuery(dbConn, query, config)
	if err != nil {
		return err
	}
	test := data["data"].([]map[string]interface{})[0]["ok"].(float64)
	if test == 1 {
		return nil
	}
	return errors.New("connection test failed")
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

func (mqe *MongoQueryEngine) GetData(dbConn *models.DBConnection, name string, limit int, offset int64, isFirstFetch bool, filter []string, sort []string, config *models.QueryConfig) (map[string]interface{}, error) {
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
	if isFirstFetch {
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
