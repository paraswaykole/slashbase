package models

import (
	"time"
)

type TeamMember struct {
	UserID    string    `gorm:"primaryKey"`
	TeamID    string    `gorm:"primaryKey"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:user_id"`
	Team Team `gorm:"foreignkey:team_id"`
}

const (
	ROLE_ADMIN     = "ADMIN"
	ROLE_DEVELOPER = "DEVELOPER"
	ROLE_ANALYST   = "ANALYST"
)

func newTeamMember(userID string, teamID string, role string) *TeamMember {
	return &TeamMember{
		UserID: userID,
		TeamID: teamID,
		Role:   role,
	}
}

func NewTeamAdmin(userID string, teamID string) *TeamMember {
	return newTeamMember(userID, teamID, ROLE_ADMIN)
}

func NewTeamDeveloper(userID string, teamID string) *TeamMember {
	return newTeamMember(userID, teamID, ROLE_DEVELOPER)
}

func NewTeamAnalyst(userID string, teamID string) *TeamMember {
	return newTeamMember(userID, teamID, ROLE_ANALYST)
}
