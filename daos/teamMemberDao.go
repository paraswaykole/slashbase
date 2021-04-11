package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

func (d TeamDao) CreateTeamMembers(teamMembers *[]models.TeamMember) error {
	result := db.GetDB().Create(teamMembers)
	return result.Error
}
