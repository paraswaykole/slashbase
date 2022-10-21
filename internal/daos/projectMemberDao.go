package daos

import (
	"errors"

	"gorm.io/gorm"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func (d ProjectDao) CreateProjectMembers(projectMembers *[]models.ProjectMember) error {
	result := db.GetDB().Create(projectMembers)
	return result.Error
}

func (d ProjectDao) CreateProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Create(projectMember)
	return result.Error
}

func (d ProjectDao) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Preload("User").Find(&projectMembers).Error
	return &projectMembers, err
}

func (d ProjectDao) DeleteAllProjectMembersInProject(projectID string) error {
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Delete(models.ProjectMember{}).Error
	return err
}

func (d ProjectDao) GetProjectMembersForUser(userID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID}).Preload("Project").Find(&projectMembers).Error
	return &projectMembers, err
}

func (d ProjectDao) FindProjectMember(projectID string, userID string) (*models.ProjectMember, bool, error) {
	var projectMember models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID, ProjectID: projectID}).Preload("Project").First(&projectMember).Error
	return &projectMember, errors.Is(err, gorm.ErrRecordNotFound), err
}

func (d ProjectDao) DeleteProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Where(models.ProjectMember{UserID: projectMember.UserID, ProjectID: projectMember.ProjectID}).Delete(&models.ProjectMember{})
	return result.Error
}
