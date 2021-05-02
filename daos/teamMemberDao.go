package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

func (d TeamDao) CreateTeamMembers(teamMembers *[]models.TeamMember) error {
	result := db.GetDB().Create(teamMembers)
	return result.Error
}

func (d TeamDao) GetTeamMembers(teamID string) (*[]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	err := db.GetDB().Where(models.TeamMember{TeamID: teamID}).Preload("User").Find(&teamMembers).Error
	return &teamMembers, err
}
