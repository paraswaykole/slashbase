package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
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
	TAB_TYPE_BLANK     = "BLANK"
	TAB_TYPE_DATAMODEL = "DATAMODEL"
	TAB_TYPE_QUERY     = "QUERY"
	TAB_TYPE_HISTORY   = "HISTORY"
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

func NewDataModelTab(dbConnID, schema, name string) *Tab {
	data := map[string]interface{}{
		"schema": schema,
		"name":   name,
	}
	dataStr, _ := json.Marshal(data)
	return newTab(TAB_TYPE_DATAMODEL, dbConnID, string(dataStr))
}

func NewQueryTab(dbConnID, queryID, query string) *Tab {
	data := map[string]interface{}{
		"queryid": queryID,
		"query":   query,
	}
	dataStr, _ := json.Marshal(data)
	return newTab(TAB_TYPE_QUERY, dbConnID, string(dataStr))
}

func NewHistoryTab(dbConnID string) *Tab {
	return newTab(TAB_TYPE_HISTORY, dbConnID, "")
}
