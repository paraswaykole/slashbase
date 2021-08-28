package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

type ProjectDao struct{}

func (d ProjectDao) CreateProject(project *models.Project) (*models.ProjectMember, error) {
	result := db.GetDB().Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	projectMember := []models.ProjectMember{*models.NewProjectAdmin(project.CreatedBy, project.ID)}
	err := d.CreateProjectMembers(&projectMember)
	if err != nil {
		return nil, result.Error
	}
	return &projectMember[0], err
}
