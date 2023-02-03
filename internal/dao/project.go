package dao

import (
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/models"
)

type projectDao struct{}

var Project projectDao

func (projectDao) CreateProject(project *models.Project) error {
	result := db.GetDB().Create(project)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (projectDao) GetProject(id string) (*models.Project, error) {
	var project models.Project
	result := db.GetDB().Where(models.Project{ID: id}).First(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (projectDao) GetAllProjects() (*[]models.Project, error) {
	var projects []models.Project
	err := db.GetDB().Model(models.Project{}).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (projectDao) GetAllProjectsCount() (int64, error) {
	var count int64
	err := db.GetDB().Model(models.Project{}).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (projectDao) DeleteProject(id string) error {
	result := db.GetDB().Where(models.Project{ID: id}).Delete(models.Project{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
