package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"index:idx_name_teamid,unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Permissions []RolePermission
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
