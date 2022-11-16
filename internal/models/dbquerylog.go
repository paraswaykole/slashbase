package models

import (
	"time"

	"github.com/google/uuid"
)

type DBQueryLog struct {
	ID             string    `gorm:"type:uuid;primaryKey"`
	Query          string    `gorm:"not null"`
	UserID         string    `gorm:"type:uuid;not null"`
	DBConnectionID string    `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:user_id"`
}

func NewQueryLog(userID string, dbConnectionID string, query string) *DBQueryLog {
	return &DBQueryLog{
		ID:             uuid.NewString(),
		Query:          query,
		UserID:         userID,
		DBConnectionID: dbConnectionID,
	}
}
