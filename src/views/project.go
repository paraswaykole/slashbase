package views

import (
	"time"

	"slashbase.com/backend/src/models"
)

type ProjectView struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CurrentMember ProjectMemberView `json:"currentMember"`
	CreatedAt     time.Time         `json:"createdAt"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

type ProjectMemberView struct {
	ID        string    `json:"id"`
	User      UserView  `json:"user"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildProject(pProject *models.Project, currentMember *models.ProjectMember) ProjectView {
	projectView := ProjectView{
		ID:            pProject.ID,
		Name:          pProject.Name,
		CurrentMember: BuildProjectMember(currentMember),
		CreatedAt:     pProject.CreatedAt,
		UpdatedAt:     pProject.UpdatedAt,
	}
	return projectView
}

func BuildProjectMember(projectMember *models.ProjectMember) ProjectMemberView {
	projectMemberView := ProjectMemberView{
		ID:        projectMember.UserID + ":" + projectMember.ProjectID,
		User:      BuildUser(&projectMember.User),
		Role:      projectMember.Role,
		CreatedAt: projectMember.CreatedAt,
		UpdatedAt: projectMember.UpdatedAt,
	}
	return projectMemberView
}
