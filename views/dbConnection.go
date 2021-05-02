package views

import (
	"time"

	"slashbase.com/backend/models"
)

type DBConnectionView struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	TeamID    string    `json:"teamId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildDBConnection(dbConn *models.DBConnection) DBConnectionView {
	dbConnView := DBConnectionView{
		ID:        dbConn.ID,
		Name:      dbConn.Name,
		Type:      dbConn.Type,
		TeamID:    dbConn.TeamID,
		CreatedAt: dbConn.CreatedAt,
		UpdatedAt: dbConn.UpdatedAt,
	}
	return dbConnView
}
