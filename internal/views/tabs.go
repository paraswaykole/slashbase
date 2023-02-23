package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/models"
)

type TabView struct {
	ID             string                 `json:"id"`
	Type           string                 `json:"type"`
	MetaData       map[string]interface{} `json:"metadata"`
	DBConnectionID string                 `json:"dbConnectionId"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
}

func BuildTabView(tab *models.Tab) *TabView {
	return &TabView{
		ID:             tab.ID,
		Type:           tab.Type,
		DBConnectionID: tab.DBConnectionID,
		MetaData:       tab.FetchMetadata(),
		CreatedAt:      tab.CreatedAt,
		UpdatedAt:      tab.UpdatedAt,
	}
}
