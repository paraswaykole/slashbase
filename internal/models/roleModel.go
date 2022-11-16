package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `gorm:"uniqueIndex"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

const (
	ROLE_ADMIN = "Admin"
)

func NewRole(name string) *Role {
	return &Role{
		ID:   uuid.NewString(),
		Name: name,
	}
}
