package mongoutils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	QUERY_FIND      = iota
	QUERY_FINDONE   = iota
	QUERY_INSERT    = iota
	QUERY_INSERTONE = iota
	QUERY_UPDATE    = iota
	QUERY_UPDATEONE = iota
	QUERY_UNKOWN    = -1
)

type MongoQuery struct {
	QueryType      int
	CollectionName string
	Filter         bson.D
}

func GetMongoQueryType(query string) *MongoQuery {
	var result MongoQuery
	tokens := strings.Split(query, ".")
	if len(tokens) == 0 || tokens[0] != "db" {
		result.QueryType = QUERY_UNKOWN
		return &result
	}
	if len(tokens) > 1 {
		result.CollectionName = tokens[1]
	}
	if len(tokens) > 2 {
		if strings.HasPrefix(tokens[2], "find(") {
			result.QueryType = QUERY_FIND
			result.Filter = bson.D{}
			// TODO: fill up filter
		} else if strings.HasPrefix(tokens[2], "findOne(") {
			result.QueryType = QUERY_FINDONE
			result.Filter = bson.D{}
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
