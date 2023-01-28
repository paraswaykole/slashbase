package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewProject(name string) (*Project, error) {

	if len(strings.TrimSpace(name)) == 0 {
		return nil, errors.New("project name cannot be empty")
	}

	return &Project{
		ID:   uuid.New().String(),
		Name: name,
	}, nil
}
