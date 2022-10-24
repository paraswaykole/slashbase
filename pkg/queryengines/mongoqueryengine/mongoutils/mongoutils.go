package mongoutils

import (
	"context"
	"log"
	"regexp"
	"strconv"
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
		rowDataMap := bsonDtoJsonMap(&rowData, &keysMap)
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
	rowDataMap := bsonDtoJsonMap(&rowData, &keysMap)
	keysList := []string{}
	for key := range keysMap {
		keysList = append(keysList, key)
	}
	return keysList, []map[string]interface{}{rowDataMap}
}

func bsonDtoJsonMap(data *bson.D, keysMap *map[string]bool) map[string]interface{} {
	dataMap := make(map[string]interface{}, len(*data))
	for _, e := range *data {
		if valueData, isTrue := e.Value.(bson.D); isTrue {
			dataMap[e.Key] = bsonDtoJsonMap(&valueData, nil)
		} else if valueData, isTrue := e.Value.(bson.A); isTrue {
			mapArray := make([]interface{}, len(valueData))
			for i, valueDataData := range valueData {
				if valueDataDataData, isTrueTrue := valueDataData.(bson.D); isTrueTrue {
					mapArray[i] = bsonDtoJsonMap(&valueDataDataData, nil)
				} else {
					mapArray[i] = valueDataData
				}
			}
			dataMap[e.Key] = mapArray
		} else {
			dataMap[e.Key] = e.Value
		}
		if keysMap != nil {
			(*keysMap)[e.Key] = true
		}
	}
	return dataMap
}

const (
	QUERY_FIND            = iota
	QUERY_FINDONE         = iota
	QUERY_INSERT          = iota
	QUERY_INSERTONE       = iota
	QUERY_UPDATE          = iota
	QUERY_UPDATEONE       = iota
	QUERY_COUNT           = iota
	QUERY_RUNCMD          = iota
	QUERY_LISTCOLLECTIONS = iota
	QUERY_UNKOWN          = -1
)

type MongoQuery struct {
	QueryType      int
	CollectionName string
	Data           interface{}
	Limit          *int64
	Skip           *int64
}

func GetMongoQueryType(query string) *MongoQuery {
	var result MongoQuery
	re := regexp.MustCompile(`(\(\{.+\}\))|(\(\[.+\]\))|(\(\d+\))|(\(\))`)
	filteredQuery := strings.ReplaceAll(query, " ", "")
	filteredQuery = re.ReplaceAllString(filteredQuery, "")
	tokens := strings.Split(filteredQuery, ".")
	if len(tokens) == 0 || tokens[0] != "db" {
		result.QueryType = QUERY_UNKOWN
		return &result
	}
	if len(tokens) > 1 {
		token := tokens[1]
		if token == "runCommand" {
			result.QueryType = QUERY_RUNCMD
			filter := findBsonOfToken(token, query)
			result.Data = filter
			return &result
		}
		if token == "getCollectionNames" {
			result.QueryType = QUERY_LISTCOLLECTIONS
			result.Data = bson.D{}
			return &result
		}
		result.CollectionName = token
	}
	if len(tokens) > 2 {
		funcName := tokens[2]
		filter := findBsonOfToken(funcName, query)
		if funcName == "find" {
			result.QueryType = QUERY_FIND
			if len(funcName) > 3 {
				for _, tkn := range tokens[3:] {
					if strings.HasPrefix(tkn, "limit(") {
						numberInterface := findBsonOfToken(tkn, query)
						if numberInterface != nil {
							number := numberInterface.(int64)
							result.Limit = &number
						}
					} else if strings.HasPrefix(tkn, "skip(") {
						numberInterface := findBsonOfToken(tkn, query)
						if numberInterface != nil {
							number := numberInterface.(int64)
							result.Skip = &number
						}
					}
				}
			}
		} else if funcName == "findOne" {
			result.QueryType = QUERY_FINDONE
		} else if funcName == "insert" {
			result.QueryType = QUERY_INSERT
		} else if funcName == "insertOne" {
			result.QueryType = QUERY_INSERTONE
		} else if funcName == "update" {
			result.QueryType = QUERY_UPDATE
		} else if funcName == "updateOne" {
			result.QueryType = QUERY_UPDATEONE
		} else if funcName == "count" {
			result.QueryType = QUERY_COUNT
		}
		result.Data = filter
	}
	return &result
}

func findBsonOfToken(tokenName string, rawQuery string) interface{} {
	re := regexp.MustCompile(tokenName + `((?:\(\{.+\}\))|(?:\(\[.+\]\))|(?:\(\d+\))|(?:\(\)))`)
	token := re.FindString(rawQuery)
	strdata := strings.Split(strings.Trim(token, ")"), "(")
	tokenData := strdata[1]
	if tokenData == "" {
		return bson.D{}
	}
	if strings.HasPrefix(tokenData, "{") && strings.HasSuffix(tokenData, "}") {
		var mapData map[string]interface{}
		err := yaml.Unmarshal([]byte(tokenData), &mapData)
		if err != nil {
			return bson.D{}
		}
		bsonData := bson.D{}
		for key, value := range mapData {
			bsonData = append(bsonData, bson.E{
				Key:   key,
				Value: value,
			})
		}
		return bsonData
	} else if strings.HasPrefix(tokenData, "[") && strings.HasPrefix(tokenData, "]") {
		var arrayData []map[string]interface{}
		err := yaml.Unmarshal([]byte(tokenData), &arrayData)
		if err != nil {
			return bson.D{}
		}
		bsonArray := make(bson.A, len(arrayData))
		for i, value := range arrayData {
			bsonData := bson.D{}
			for key, value := range value {
				bsonData = append(bsonData, bson.E{
					Key:   key,
					Value: value,
				})
			}
			bsonArray[i] = bsonData
		}
		return bsonArray
	} else {
		strdata := strings.Split(strings.Trim(token, ")"), "(")
		if strdata[1] == "" {
			return nil
		}
		number, err := strconv.ParseInt(strdata[1], 10, 64)
		if err != nil {
			return nil
		}
		return number
	}
}
