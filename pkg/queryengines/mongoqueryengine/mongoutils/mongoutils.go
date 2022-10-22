package mongoutils

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func MongoCursorToJson(cur *mongo.Cursor) ([]string, []map[string]interface{}) {
	keysMap := map[string]bool{}
	resultData := make([]map[string]interface{}, 0)
	for cur.Next(context.Background()) {
		var rowData bson.D
		err := cur.Decode(&rowData)
		if err != nil {
			log.Fatal(err)
		}
		rowDataMap := make(map[string]interface{}, len(rowData))
		for _, e := range rowData {
			rowDataMap[e.Key] = e.Value
			keysMap[e.Key] = true
		}
		resultData = append(resultData, rowDataMap)
	}
	keysList := []string{}
	for key := range keysMap {
		keysList = append(keysList, key)
	}
	return keysList, resultData
}

func MongoSingleResultToJson(result *mongo.SingleResult) ([]string, []map[string]interface{}) {
	keysMap := map[string]bool{}
	var rowData bson.D
	err := result.Decode(&rowData)
	if err != nil {
		return nil, nil
	}
	rowDataMap := make(map[string]interface{}, len(rowData))
	for _, e := range rowData {
		rowDataMap[e.Key] = e.Value
		keysMap[e.Key] = true
	}
	keysList := []string{}
	for key := range keysMap {
		keysList = append(keysList, key)
	}
	return keysList, []map[string]interface{}{rowDataMap}
}

const (
	QUERY_FIND            = iota
	QUERY_FINDONE         = iota
	QUERY_INSERT          = iota
	QUERY_INSERTONE       = iota
	QUERY_UPDATE          = iota
	QUERY_UPDATEONE       = iota
	QUERY_RUNCMD          = iota
	QUERY_LISTCOLLECTIONS = iota
	QUERY_UNKOWN          = -1
)

type MongoQuery struct {
	QueryType      int
	CollectionName string
	Data           bson.D
}

func GetMongoQueryType(query string) *MongoQuery {
	var result MongoQuery
	tokens := strings.Split(query, ".")
	if len(tokens) == 0 || tokens[0] != "db" {
		result.QueryType = QUERY_UNKOWN
		return &result
	}
	if len(tokens) > 1 {
		token := tokens[1]
		if strings.HasPrefix(token, "runCommand(") {
			result.QueryType = QUERY_RUNCMD
			_, filter := splitToken(token)
			result.Data = filter
			return &result
		}
		if strings.HasPrefix(token, "getCollectionNames(") {
			result.QueryType = QUERY_LISTCOLLECTIONS
			result.Data = bson.D{}
			return &result
		}
		result.CollectionName = token
	}
	if len(tokens) > 2 {
		if strings.HasPrefix(tokens[2], "find(") {
			result.QueryType = QUERY_FIND
			result.Data = bson.D{}
			// TODO: fill up filter
		} else if strings.HasPrefix(tokens[2], "findOne(") {
			result.QueryType = QUERY_FINDONE
			result.Data = bson.D{}
			// TODO: fill up filter
		} else if strings.HasPrefix(tokens[2], "insert(") {
			result.QueryType = QUERY_INSERT
		} else if strings.HasPrefix(tokens[2], "insertOne(") {
			result.QueryType = QUERY_INSERTONE
		} else if strings.HasPrefix(tokens[2], "update(") {
			result.QueryType = QUERY_UPDATE
		} else if strings.HasPrefix(tokens[2], "updateOne(") {
			result.QueryType = QUERY_UPDATEONE
		}
	}
	return &result
}

func splitToken(token string) (string, bson.D) {
	strdata := strings.Split(strings.Trim(token, ")"), "(")
	bsonData := bson.D{}
	if strdata[1] == "" {
		return strdata[0], bsonData
	}
	var mapData map[string]interface{}
	err := yaml.Unmarshal([]byte(strdata[1]), &mapData)
	if err != nil {
		return strdata[0], bsonData
	}
	for key, value := range mapData {
		bsonData = append(bsonData, bson.E{
			Key:   key,
			Value: value,
		})
	}
	return strdata[0], bsonData
}
