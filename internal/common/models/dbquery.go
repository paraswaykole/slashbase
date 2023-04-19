package models

import (
	"time"

	"github.com/google/uuid"
)

type DBQuery struct {
	ID             string    `gorm:"type:uuid;primaryKey"`
	Name           string    `gorm:"not null"`
	Query          string    `gorm:"not null"`
	DBConnectionID string    `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	DBConnection DBConnection `gorm:"foreignkey:db_connection_id"`
}

func NewQuery(name string, query string, dbConnectionID string) *DBQuery {
	return &DBQuery{
		ID:             uuid.New().String(),
		Name:           name,
		Query:          query,
		DBConnectionID: dbConnectionID,
	}
}
