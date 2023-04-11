package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"index:idx_name_teamid,unique"`
	TeamID    string    `gorm:"index:idx_name_teamid,unique;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Permissions []RolePermission
}

const (
	ROLE_ADMIN = "Admin"
)

func NewRole(name string, teamID string) *Role {
	return &Role{
		ID:     uuid.NewString(),
		Name:   name,
		TeamID: teamID,
	}
}
