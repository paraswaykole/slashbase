package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/common/models"
)

type DBConnectionView struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	ProjectID string    `json:"projectId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildDBConnection(dbConn *models.DBConnection) DBConnectionView {
	dbConnView := DBConnectionView{
		ID:        dbConn.ID,
		Name:      dbConn.Name,
		Type:      dbConn.Type,
		ProjectID: dbConn.ProjectID,
		CreatedAt: dbConn.CreatedAt,
		UpdatedAt: dbConn.UpdatedAt,
	}
	return dbConnView
}
