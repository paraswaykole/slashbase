package views

import (
	"time"

	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/models"
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
	Role      RoleView  `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildProject(pProject *common.Project, currentMember *models.ProjectMember) ProjectView {
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
		Role:      BuildRole(&projectMember.Role),
		CreatedAt: projectMember.CreatedAt,
		UpdatedAt: projectMember.UpdatedAt,
	}
	return projectMemberView
}
