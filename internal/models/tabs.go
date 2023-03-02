package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/db"
)

type Tab struct {
	ID             string `gorm:"type:uuid;primaryKey"`
	Type           string `gorm:"not null"`
	DBConnectionID string `gorm:"type:uuid;not null"`
	MetaData       string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	DBConnection DBConnection `gorm:"foreignkey:db_connection_id"`
}

const (
	TAB_TYPE_BLANK   = "BLANK"
	TAB_TYPE_DATA    = "DATA"
	TAB_TYPE_MODEL   = "MODEL"
	TAB_TYPE_QUERY   = "QUERY"
	TAB_TYPE_HISTORY = "HISTORY"
	TAB_TYPE_CONSOLE = "CONSOLE"
)

func newTab(ttype, dbConnID, metaData string) *Tab {
	return &Tab{
		ID:             uuid.New().String(),
		Type:           ttype,
		DBConnectionID: dbConnID,
		MetaData:       metaData,
	}
}

func NewBlankTab(dbConnID string) *Tab {
	return newTab(TAB_TYPE_BLANK, dbConnID, "")
}

func NewDataTab(dbConnID, schema, name string) *Tab {
	data := map[string]interface{}{
		"schema": schema,
		"name":   name,
	}
	dataStr, _ := json.Marshal(data)
	return newTab(TAB_TYPE_DATA, dbConnID, string(dataStr))
}

func NewModelTab(dbConnID, schema, name string) *Tab {
	data := map[string]interface{}{
		"schema": schema,
		"name":   name,
	}
	dataStr, _ := json.Marshal(data)
	return newTab(TAB_TYPE_MODEL, dbConnID, string(dataStr))
}

func NewQueryTab(dbConnID, queryID, query string) *Tab {
	data := map[string]interface{}{
		"queryId": queryID,
		"query":   query,
	}
	dataStr, _ := json.Marshal(data)
	return newTab(TAB_TYPE_QUERY, dbConnID, string(dataStr))
}

func NewHistoryTab(dbConnID string) *Tab {
	return newTab(TAB_TYPE_HISTORY, dbConnID, "")
}

func NewConsoleTab(dbConnID string) *Tab {
	return newTab(TAB_TYPE_CONSOLE, dbConnID, "")
}

func (t *Tab) FetchMetadata() map[string]interface{} {
	var metadata map[string]interface{}
	err := json.Unmarshal([]byte(t.MetaData), &metadata)
	if err != nil {
		metadata = map[string]interface{}{}
	}
	if t.Type == TAB_TYPE_QUERY {
		queryID := metadata["queryId"].(string)
		var dbQuery DBQuery
		err := db.GetDB().Where(DBQuery{ID: queryID}).First(&dbQuery).Error
		if err == nil {
			metadata["queryName"] = dbQuery.Name
		}
	}
	return metadata
}
