package models

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/common/models"
)

type ProjectMember struct {
	UserID    string `gorm:"primaryKey"`
	ProjectID string `gorm:"primaryKey"`
	RoleID    string
	IsCreator bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User    User           `gorm:"foreignkey:user_id"`
	Project models.Project `gorm:"foreignkey:project_id"`
	Role    Role           `gorm:"foreignkey:role_id;constraint:OnDelete:SET NULL;"`
}

func NewProjectMember(userID string, projectID string, roleID string) *ProjectMember {
	return &ProjectMember{
		UserID:    userID,
		ProjectID: projectID,
		RoleID:    roleID,
	}
}
