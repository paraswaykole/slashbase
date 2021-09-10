package views

import "slashbase.com/backend/models"

type DBQueryView struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Query          string `json:"query"`
	DBConnectionID string `json:"dbConnectionId"`
}

func BuildDBQueryView(query *models.DBQuery) *DBQueryView {
	return &DBQueryView{
		ID:             query.ID,
		Name:           query.Name,
		Query:          query.Query,
		DBConnectionID: query.DBConnectionID,
	}
}
