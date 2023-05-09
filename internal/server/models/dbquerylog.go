package models

import (
	"github.com/slashbaseide/slashbase/internal/common/models"
)

type DBQueryLog struct {
	models.DBQueryLog
	UserID string `gorm:"primaryKey"`

	User User `gorm:"foreignkey:user_id"`
}

func NewUserQueryLog(queryLog models.DBQueryLog, userID string) *DBQueryLog {
	return &DBQueryLog{
		DBQueryLog: queryLog,
		UserID:     userID,
	}
}
