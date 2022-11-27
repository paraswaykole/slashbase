package dao

import (
	"errors"

	"gorm.io/gorm"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func (projectDao) CreateProjectMembers(projectMembers *[]models.ProjectMember) error {
	result := db.GetDB().Create(projectMembers)
	return result.Error
}

func (projectDao) CreateProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Create(projectMember)
	return result.Error
}

func (projectDao) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Preload("User").Preload("Role").Find(&projectMembers).Error
	return &projectMembers, err
}

func (projectDao) DeleteAllProjectMembersInProject(projectID string) error {
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Delete(models.ProjectMember{}).Error
	return err
}

func (projectDao) GetProjectMembersForUser(userID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID}).Preload("Project").Preload("Role").Find(&projectMembers).Error
	return &projectMembers, err
}

func (projectDao) FindProjectMember(projectID string, userID string) (*models.ProjectMember, bool, error) {
	var projectMember models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID, ProjectID: projectID}).Preload("Project").Preload("Role").First(&projectMember).Error
	return &projectMember, errors.Is(err, gorm.ErrRecordNotFound), err
}

func (projectDao) DeleteProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Where(models.ProjectMember{UserID: projectMember.UserID, ProjectID: projectMember.ProjectID}).Delete(&models.ProjectMember{})
	return result.Error
}
