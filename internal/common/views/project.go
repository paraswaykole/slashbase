package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/common/models"
)

type ProjectView struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildProject(pProject *models.Project) ProjectView {
	projectView := ProjectView{
		ID:        pProject.ID,
		Name:      pProject.Name,
		CreatedAt: pProject.CreatedAt,
		UpdatedAt: pProject.UpdatedAt,
	}
	return projectView
}
