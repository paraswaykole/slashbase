package daos

import (
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
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

func (d ProjectDao) GetProject(id string) (*models.Project, error) {
	var project models.Project
	result := db.GetDB().Where(models.Project{ID: id}).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (d ProjectDao) DeleteProject(id string) error {
	result := db.GetDB().Where(models.Project{ID: id}).Delete(models.Project{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
