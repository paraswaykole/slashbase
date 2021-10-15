package daos

import (
	"slashbase.com/backend/src/db"
	"slashbase.com/backend/src/models"
)

type ProjectDao struct{}

func (d ProjectDao) CreateProject(project *models.Project) (*models.ProjectMember, error) {
	result := db.GetDB().Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	projectMember := models.NewProjectAdmin(project.CreatedBy, project.ID)
	err := d.CreateProjectMember(projectMember)
	if err != nil {
		return nil, result.Error
	}
	return projectMember, err
}
