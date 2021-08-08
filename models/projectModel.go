package models

import (
	"time"
)

type Project struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string    `gorm:"not null"`
	CreatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	CreatedByUser User `gorm:"foreignkey:created_by"`
}

func NewProject(createdBy *User, name string) *Project {
	return &Project{
		Name:      name,
		CreatedBy: createdBy.ID,
	}
}
