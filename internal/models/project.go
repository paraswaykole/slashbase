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
}

func NewProject(name string) *Project {
	return &Project{
		ID:   uuid.NewString(),
		Name: name,
	}
}
