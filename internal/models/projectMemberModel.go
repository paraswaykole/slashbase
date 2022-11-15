package models

import (
	"time"
)

type ProjectMember struct {
	UserID    string `gorm:"primaryKey"`
	ProjectID string `gorm:"primaryKey"`
	RoleID    string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User    User    `gorm:"foreignkey:user_id"`
	Project Project `gorm:"foreignkey:project_id"`
	Role    Role    `gorm:"foreignkey:role_id"`
}

func NewProjectMember(userID string, projectID string, roleID string) *ProjectMember {
	return &ProjectMember{
		UserID:    userID,
		ProjectID: projectID,
		RoleID:    roleID,
	}
}
