package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/server/models"
)

type DBQueryLogView struct {
	ID             string    `json:"id"`
	Query          string    `json:"query"`
	DBConnectionID string    `json:"dbConnectionId"`
	User           UserView  `json:"user"`
	CreatedAt      time.Time `json:"createdAt"`
}

func BuildDBQueryLogView(log *models.DBQueryLog) *DBQueryLogView {
	return &DBQueryLogView{
		ID:             log.ID,
		Query:          log.Query,
		DBConnectionID: log.DBConnectionID,
		User:           BuildUser(&log.User),
		CreatedAt:      log.CreatedAt,
	}
}
