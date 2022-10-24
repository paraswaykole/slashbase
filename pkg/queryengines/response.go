package queryengines

import "slashbase.com/backend/internal/models"

type AddDataResponse struct {
	NewID string `json:"newId"`
}

func BuildAddDataResponse(dbConn *models.DBConnection, queryData map[string]interface{}) *AddDataResponse {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		view := AddDataResponse{
			NewID: queryData["ctid"].(string),
		}
		return &view
	} else if dbConn.Type == models.DBTYPE_MONGO {
		view := AddDataResponse{
			NewID: queryData["insertedId"].(string),
		}
		return &view
	}
	return nil
}
