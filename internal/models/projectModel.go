package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedBy string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	CreatedByUser User `gorm:"foreignkey:created_by"`
}

func NewProject(createdBy *User, name string) *Project {
	return &Project{
		ID:        uuid.NewString(),
		Name:      name,
		CreatedBy: createdBy.ID,
	}
}
