package models

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	RoleID    string `gorm:"index:idx_roleid_name,unique"`
	Name      string `gorm:"index:idx_roleid_name,unique"`
	Value     bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Role Role `gorm:"foreignkey:role_id;constraint:OnDelete:SET NULL;"`
}

const (
	ROLE_PERMISSION_NAME_READ_ONLY = "READ_ONLY"
)

func NewRolePermission(roleID string, name string, value bool) *RolePermission {
	return &RolePermission{
		ID:     uuid.NewString(),
		RoleID: roleID,
		Name:   name,
		Value:  value,
	}
}
