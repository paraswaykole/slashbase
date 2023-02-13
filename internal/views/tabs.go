package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/models"
)

type TabView struct {
	Type           string    `json:"type"`
	DBConnectionID string    `json:"dbConnectionId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func BuildTabView(tab *models.Tab) *TabView {
	return &TabView{
		Type:           tab.Type,
		DBConnectionID: tab.DBConnectionID,
		CreatedAt:      tab.CreatedAt,
		UpdatedAt:      tab.UpdatedAt,
	}
}
