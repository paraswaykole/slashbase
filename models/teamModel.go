package models

import (
	"time"
)

type Team struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `gorm:"not null"`
	CreatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	CreatedByUser User `gorm:"foreignkey:created_by"`
}

func NewTeam(createdBy *User, name string) *Team {
	return &Team{
		Name:      name,
		CreatedBy: createdBy.ID,
	}
}
