package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

type TeamDao struct{}

func (d TeamDao) CreateTeam(team *models.Team) error {
	result := db.GetDB().Create(team)
	if result.Error != nil {
		return result.Error
	}
	teamMember := []models.TeamMember{*models.NewTeamAdmin(team.CreatedBy, team.ID)}
	err := d.CreateTeamMembers(&teamMember)
	return err
}
