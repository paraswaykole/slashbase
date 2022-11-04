package models

import (
	"time"

	"github.com/google/uuid"
)

type DBQuery struct {
	ID             string    `gorm:"type:uuid;primaryKey"`
	Name           string    `gorm:"not null"`
	Query          string    `gorm:"not null"`
	CreatedBy      string    `gorm:"not null"`
	DBConnectionID string    `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	CreatedByUser User         `gorm:"foreignkey:created_by"`
	DBConnection  DBConnection `gorm:"foreignkey:db_connection_id"`
}

func NewQuery(createdBy *User, name string, query string, dbConnectionID string) *DBQuery {
	return &DBQuery{
		ID:             uuid.NewString(),
		Name:           name,
		Query:          query,
		DBConnectionID: dbConnectionID,
		CreatedBy:      createdBy.ID,
	}
}
