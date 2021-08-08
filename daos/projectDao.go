package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

type ProjectDao struct{}

func (d ProjectDao) CreateProject(project *models.Project) error {
	result := db.GetDB().Create(project)
	if result.Error != nil {
		return result.Error
	}
	projectMember := []models.ProjectMember{*models.NewProjectAdmin(project.CreatedBy, project.ID)}
	err := d.CreateProjectMembers(&projectMember)
	return err
}
