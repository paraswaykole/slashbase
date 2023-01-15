package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddDataResponse struct {
	NewID string                 `json:"newId"`
	Data  map[string]interface{} `json:"data"`
}

func BuildAddDataResponse(dbConn *DBConnection, queryData map[string]interface{}) *AddDataResponse {
	if dbConn.Type == DBTYPE_POSTGRES {
		ctid := queryData["ctid"].(string)
		delete(queryData, "ctid")
		view := AddDataResponse{
			NewID: ctid,
			Data:  queryData["data"].(map[string]interface{}),
		}
		return &view
	} else if dbConn.Type == DBTYPE_MONGO {
		view := AddDataResponse{
			NewID: queryData["insertedId"].(primitive.ObjectID).Hex(),
		}
		return &view
	} else if dbConn.Type == DBTYPE_MYSQL {
		view := AddDataResponse{
			Data: queryData["data"].(map[string]interface{}),
		}
		return &view
	}
	return nil
}
