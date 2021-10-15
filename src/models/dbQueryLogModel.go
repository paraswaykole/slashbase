package models

import (
	"time"
)

type DBQueryLog struct {
	ID             string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Query          string    `gorm:"not null"`
	UserID         string    `gorm:"type:uuid;not null"`
	DBConnectionID string    `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:user_id"`
}

func NewQueryLog(userID string, dbConnectionID string, query string) *DBQueryLog {
	return &DBQueryLog{
		Query:          query,
		UserID:         userID,
		DBConnectionID: dbConnectionID,
	}
}
