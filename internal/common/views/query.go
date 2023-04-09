package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/common/models"
)

type DBQueryView struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Query          string `json:"query"`
	DBConnectionID string `json:"dbConnectionId"`
}

type DBQueryLogView struct {
	ID             string    `json:"id"`
	Query          string    `json:"query"`
	DBConnectionID string    `json:"dbConnectionId"`
	CreatedAt      time.Time `json:"createdAt"`
}

func BuildDBQueryView(query *models.DBQuery) *DBQueryView {
	return &DBQueryView{
		ID:             query.ID,
		Name:           query.Name,
		Query:          query.Query,
		DBConnectionID: query.DBConnectionID,
	}
}

func BuildDBQueryLogView(log *models.DBQueryLog) *DBQueryLogView {
	return &DBQueryLogView{
		ID:             log.ID,
		Query:          log.Query,
		DBConnectionID: log.DBConnectionID,
		CreatedAt:      log.CreatedAt,
	}
}
