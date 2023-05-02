package models

import (
	"time"

	"github.com/google/uuid"
)

type DBQueryLog struct {
	ID             string    `gorm:"type:uuid;primaryKey"`
	Query          string    `gorm:"not null"`
	DBConnectionID string    `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}

func NewQueryLog(dbConnectionID string, query string) *DBQueryLog {
	return &DBQueryLog{
		ID:             uuid.New().String(),
		Query:          query,
		DBConnectionID: dbConnectionID,
	}
}
