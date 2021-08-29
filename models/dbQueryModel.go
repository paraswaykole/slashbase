package models

import (
	"time"
)

type DBQuery struct {
	ID             string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name           string    `gorm:"not null"`
	Query          string    `gorm:"not null"`
	CreatedBy      string    `gorm:"not null"`
	DBConnectionID string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	CreatedByUser User         `gorm:"foreignkey:created_by"`
	DBConnection  DBConnection `gorm:"foreignkey:db_connection_id"`
}

func NewQuery(createdBy *User, name string, query string, dbConnectionID string) *DBQuery {
	return &DBQuery{
		Name:           name,
		Query:          query,
		DBConnectionID: dbConnectionID,
		CreatedBy:      createdBy.ID,
	}
}
