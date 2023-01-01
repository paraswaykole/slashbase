package queryengines

import (
	"github.com/slashbaseide/slashbase/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddDataResponse struct {
	NewID string                 `json:"newId"`
	Data  map[string]interface{} `json:"data"`
}

func BuildAddDataResponse(dbConn *models.DBConnection, queryData map[string]interface{}) *AddDataResponse {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		ctid := queryData["ctid"].(string)
		delete(queryData, "ctid")
		view := AddDataResponse{
			NewID: ctid,
			Data:  queryData["data"].(map[string]interface{}),
		}
		return &view
	} else if dbConn.Type == models.DBTYPE_MONGO {
		view := AddDataResponse{
			NewID: queryData["insertedId"].(primitive.ObjectID).Hex(),
		}
		return &view
	}
	return nil
}
