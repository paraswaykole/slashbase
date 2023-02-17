package views

import (
	"encoding/json"
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
	var metadata map[string]interface{}
	err := json.Unmarshal([]byte(tab.MetaData), &metadata)
	if err != nil {
		metadata = map[string]interface{}{}
	}
	return &TabView{
		ID:             tab.ID,
		Type:           tab.Type,
		DBConnectionID: tab.DBConnectionID,
		MetaData:       metadata,
		CreatedAt:      tab.CreatedAt,
		UpdatedAt:      tab.UpdatedAt,
	}
}
