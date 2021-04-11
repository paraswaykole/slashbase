package views

import (
	"time"

	"slashbase.com/backend/models"
)

type TeamView struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildTeam(pTeam *models.Team) TeamView {
	teamView := TeamView{
		ID:        pTeam.ID,
		Name:      pTeam.Name,
		CreatedAt: pTeam.CreatedAt,
		UpdatedAt: pTeam.UpdatedAt,
	}
	return teamView
}
