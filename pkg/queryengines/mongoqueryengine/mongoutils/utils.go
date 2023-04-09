package mongoutils

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	QUERY_DELETEONE       = iota
	QUERY_DELETEMANY      = iota
	QUERY_UPDATEONE       = iota
	QUERY_UPDATEMANY      = iota
	QUERY_REPLACEONE      = iota
	QUERY_COUNT           = iota
	QUERY_AGGREGATE       = iota
	QUERY_GETINDEXES      = iota
	QUERY_CREATEINDEX     = iota
	QUERY_DROP            = iota
	QUERY_DROPINDEX       = iota
	QUERY_RUNCMD          = iota
	QUERY_LISTCOLLECTIONS = iota
	QUERY_UNKOWN          = -1
)

type MongoQuery struct {
	QueryType      int
	CollectionName string
	Args           []interface{}
	Limit          *int64
	Skip           *int64
	Sort           interface{}
}

func IsQueryTypeRead(query *MongoQuery) bool {
	if utils.ContainsInt([]int{QUERY_FIND, QUERY_FINDONE, QUERY_GETINDEXES, QUERY_LISTCOLLECTIONS, QUERY_COUNT}, query.QueryType) {
		return true
	}
	if query.QueryType == QUERY_AGGREGATE {
		if ops, ok := query.Args[0].(bson.A); ok {
			for _, op := range ops {
				opmap := op.(bson.D).Map()
				if _, exists := opmap["$out"]; exists {
					return false
				}
				if _, exists := opmap["$merge"]; exists {
					return false
				}
			}
		}
		return true
	}
	return false
}

func GetMongoQueryType(query string) *MongoQuery {
	var result MongoQuery = MongoQuery{
		QueryType: QUERY_UNKOWN,
	}
	tokenNames, arguments, _ := JsToTokensLexer(query)
	if len(tokenNames) == 0 || tokenNames[0] != "db" {
		result.QueryType = QUERY_UNKOWN
		return &result
	}
	if len(tokenNames) > 1 {
		argsStrList := arguments[1]
		tokenName := tokenNames[1]
		if tokenName == "runCommand" {
			result.QueryType = QUERY_RUNCMD
			args := parseTokenArgs(argsStrList)
			result.Args = args
			return &result
		}
		if tokenName == "getCollectionNames" {
			result.QueryType = QUERY_LISTCOLLECTIONS
			result.Args = []interface{}{bson.D{}}
			return &result
		}
		result.CollectionName = tokenName
	}
	if len(tokenNames) > 2 {
		funcName := tokenNames[2]
		argsStrList := arguments[2]
		args := parseTokenArgs(argsStrList)
		if funcName == "find" {
			result.QueryType = QUERY_FIND
			if len(tokenNames) > 3 {
				for i, fName := range tokenNames[3:] {
					fArg := arguments[3+i]
					if fName == "limit" {
						numberInterface := parseTokenArgs(fArg)
						if numberInterface != nil {
							number := numberInterface[0].(int64)
							result.Limit = &number
						}
					} else if fName == "skip" {
						numberInterface := parseTokenArgs(fArg)
						if numberInterface != nil {
							number := numberInterface[0].(int64)
							result.Skip = &number
						}
					} else if fName == "sort" {
						dataInterface := parseTokenArgs(fArg)
						result.Sort = dataInterface[0]
					}
				}
			}
		} else if funcName == "findOne" {
			result.QueryType = QUERY_FINDONE
		} else if funcName == "insert" {
			result.QueryType = QUERY_INSERT
		} else if funcName == "insertOne" {
			result.QueryType = QUERY_INSERTONE
		} else if funcName == "deleteOne" {
			result.QueryType = QUERY_DELETEONE
		} else if funcName == "deleteMany" {
			result.QueryType = QUERY_DELETEMANY
		} else if funcName == "updateOne" {
			result.QueryType = QUERY_UPDATEONE
		} else if funcName == "updateMany" {
			result.QueryType = QUERY_UPDATEMANY
		} else if funcName == "replaceOne" {
			result.QueryType = QUERY_REPLACEONE
		} else if funcName == "count" {
			result.QueryType = QUERY_COUNT
			if len(args) > 1 {
				options := args[1].(bson.D)
				for key, value := range options.Map() {
					if key == "limit" {
						val := int64(value.(int))
						result.Limit = &val
					} else if key == "skip" {
						val := int64(value.(int))
						result.Skip = &val
					}
				}
			}
		} else if funcName == "aggregate" {
			result.QueryType = QUERY_AGGREGATE
		} else if funcName == "getIndexes" {
			result.QueryType = QUERY_GETINDEXES
		} else if funcName == "dropIndex" {
			result.QueryType = QUERY_DROPINDEX
		} else if funcName == "drop" {
			result.QueryType = QUERY_DROP
		} else if funcName == "createIndex" {
			result.QueryType = QUERY_CREATEINDEX
		}
		result.Args = args
	}
	return &result
}

func parseTokenArgs(argsData []string) []interface{} {
	if len(argsData) == 0 {
		return []interface{}{nil}
	}
	if len(argsData) == 1 && argsData[0] == "" {
		return []interface{}{bson.D{}}
	}
	finalArgs := []interface{}{}
	for _, nArg := range argsData {
		arg := strings.TrimSpace(nArg)
		if strings.HasPrefix(arg, "{") && strings.HasSuffix(arg, "}") {
			var mapData map[string]interface{}
			err := yaml.Unmarshal([]byte(arg), &mapData)
			if err != nil {
				continue
			}
			finalArgs = append(finalArgs, mapToBsonD(&mapData))
		} else if strings.HasPrefix(arg, "[") && strings.HasSuffix(arg, "]") {
			var arrayData []map[string]interface{}
			err := yaml.Unmarshal([]byte(arg), &arrayData)
			if err != nil {
				continue
			}
			bsonArray := make(bson.A, len(arrayData))
			for i, value := range arrayData {
				bsonArray[i] = mapToBsonD(&value)
			}
			finalArgs = append(finalArgs, bsonArray)
		} else if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
			arg = strings.TrimPrefix(arg, "\"")
			arg = strings.TrimSuffix(arg, "\"")
			finalArgs = append(finalArgs, arg)
		} else if number, err := strconv.ParseInt(arg, 10, 64); err == nil {
			finalArgs = append(finalArgs, number)
		} else {
			finalArgs = append(finalArgs, arg)
		}
	}
	return finalArgs
}

func JsToTokensLexer(query string) (tokenNames []string, args [][]string, startsAt []int) {
	myReader := strings.NewReader(query)
	l := js.NewLexer(parse.NewInput(myReader))
	charLens := 0
	argStr := ""
	argsStr := []string{}
	isParenOpen := []bool{}
	isBracketBraceOpen := []bool{}
	isLastDotToken := true
	for {
		tt, text := l.Next()
		if len(isParenOpen) > 0 {
			addText := true
			if tt == js.CommaToken && len(isBracketBraceOpen) == 0 {
				addText = false
			}
			if len(isParenOpen) == 1 && tt == js.CloseParenToken {
				addText = false
			}
			if addText {
				argStr = argStr + string(text)
			}
		}
		if tt == js.IdentifierToken && isLastDotToken {
			tokenNames = append(tokenNames, string(text))
			startsAt = append(startsAt, charLens)
			isLastDotToken = false
		}
		if tt == js.DotToken {
			if len(tokenNames) != len(args) && len(isParenOpen) == 0 {
				args = append(args, nil)
			}
			isLastDotToken = true
		}
		if tt == js.CommaToken && len(isBracketBraceOpen) == 0 {
			argsStr = append(argsStr, argStr)
			argStr = ""
		}
		if tt == js.OpenParenToken {
			isParenOpen = append(isParenOpen, true)
		}
		if tt == js.OpenBracketToken || tt == js.OpenBraceToken {
			isBracketBraceOpen = append(isBracketBraceOpen, true)
		}
		if tt == js.CloseBracketToken || tt == js.CloseBraceToken {
			isBracketBraceOpen = isBracketBraceOpen[:len(isBracketBraceOpen)-1]
		}
		if tt == js.CloseParenToken {
			if len(isParenOpen) == 1 {
				argsStr = append(argsStr, argStr)
				args = append(args, argsStr)
				argStr = ""
				argsStr = []string{}
			}
			isParenOpen = isParenOpen[:len(isParenOpen)-1]
		}
		if tt == js.ErrorToken {
			break
		}
		charLens = charLens + len(text)
	}
	return
}

func mapToBsonD(data *map[string]interface{}) bson.D {
	bsonData := bson.D{}
	for key, value := range *data {
		var valueItr interface{}
		if mapItr, isTrue := value.(map[interface{}]interface{}); isTrue {
			mapData := map[string]interface{}{}
			for mapKey, mapValue := range mapItr {
				if mKey, isTrueTrue := mapKey.(string); isTrueTrue {
					mapData[mKey] = mapValue
				}
			}
			valueItr = mapToBsonD(&mapData)
		}
		if arrayItr, isTrue := value.([]interface{}); isTrue {
			var array bson.A
			for _, item := range arrayItr {
				if mapItr, isTrue := value.(map[string]interface{}); isTrue {
					array = append(array, mapToBsonD(&mapItr))
				} else if str, isTrue := item.(string); isTrue {
					if oID := stringToObjectID(str); oID != nil {
						array = append(array, *oID)
					} else {
						array = append(array, str)
					}
				} else {
					array = append(array, item)
				}
			}
			valueItr = array
		}
		if str, isTrue := value.(string); isTrue {
			if oID := stringToObjectID(str); oID != nil {
				valueItr = *oID
			} else {
				valueItr = str
			}
		}
		if valueItr == nil {
			valueItr = value
		}
		bsonData = append(bsonData, bson.E{
			Key:   key,
			Value: valueItr,
		})
	}
	return bsonData
}

func stringToObjectID(str string) *primitive.ObjectID {
	if strings.HasPrefix(str, "ObjectId(\"") && strings.HasSuffix(str, "\")") {
		hex := strings.TrimPrefix(str, "ObjectId(\"")
		hex = strings.TrimSuffix(hex, "\")")
		objectID, _ := primitive.ObjectIDFromHex(hex)
		return &objectID
	}
	return nil
}

func AnalyseFieldsSchema(keys []string, sampleData []map[string]interface{}) []map[string]interface{} {
	fields := []map[string]interface{}{}
	fieldType := map[string]map[string]bool{}
	for _, d := range sampleData {
		for key, value := range d {
			types := fieldType[key]
			if types == nil {
				types = map[string]bool{}
			}
			if value == nil {
				types["null"] = true
			} else if _, isTrue := value.(string); isTrue {
				types["string"] = true
			} else if _, isTrue := value.(int32); isTrue {
				types["int32"] = true
			} else if _, isTrue := value.(int64); isTrue {
				types["int64"] = true
			} else if _, isTrue := value.(float32); isTrue {
				types["float32"] = true
			} else if _, isTrue := value.(float64); isTrue {
				types["float64"] = true
			} else if _, isTrue := value.(primitive.ObjectID); isTrue {
				types["objectid"] = true
			} else if _, isTrue := value.(primitive.DateTime); isTrue {
				types["datetime"] = true
			} else if _, isTrue := value.([]interface{}); isTrue {
				types["array"] = true
			} else {
				types["object"] = true
			}
			fieldType[key] = types
		}
	}
	for _, key := range keys {
		types := []string{}
		for key := range fieldType[key] {
			types = append(types, key)
		}
		field := map[string]interface{}{
			"name":       key,
			"types":      strings.Join(types, ", "),
			"isNullable": utils.ContainsString(types, "null"),
			"isPrimary":  key == "_id",
		}
		fields = append(fields, field)
	}
	return fields
}

func GetCollectionIndexes(indexesData []map[string]interface{}) []map[string]interface{} {
	extractKey := func(d map[string]interface{}) interface{} {
		data, err := json.Marshal(d)
		if err != nil {
			return nil
		}
		return string(data)
	}
	indexes := []map[string]interface{}{}
	for _, index := range indexesData {
		indexes = append(indexes, map[string]interface{}{
			"name": index["name"],
			"key":  extractKey(index["key"].(map[string]interface{})),
		})
	}
	return indexes
}
