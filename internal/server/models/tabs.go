package models

import (
	"github.com/slashbaseide/slashbase/internal/common/models"
)

type Tab struct {
	models.Tab
	UserID string `gorm:"primaryKey"`

	User User `gorm:"foreignkey:user_id"`
}

func NewUserTab(tab *models.Tab, userID string) *Tab {
	return &Tab{
		UserID: userID,
		Tab:    *tab,
	}
}
