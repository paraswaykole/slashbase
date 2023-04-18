package dao

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"gorm.io/gorm"
)

type projectMemberDao struct{}

var ProjectMember projectMemberDao

func (projectMemberDao) CreateProjectMembers(projectMembers *[]models.ProjectMember) error {
	result := db.GetDB().Create(projectMembers)
	return result.Error
}

func (projectMemberDao) CreateProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Create(projectMember)
	return result.Error
}

func (projectMemberDao) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Preload("User").Preload("Role").Find(&projectMembers).Error
	return &projectMembers, err
}

func (projectMemberDao) DeleteAllProjectMembersInProject(projectID string) error {
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Delete(models.ProjectMember{}).Error
	return err
}

func (projectMemberDao) GetProjectMembersForUser(userID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID}).Preload("Project").Preload("Role.Permissions").Find(&projectMembers).Error
	return &projectMembers, err
}

func (projectMemberDao) FindProjectMember(projectID string, userID string) (*models.ProjectMember, bool, error) {
	var projectMember models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID, ProjectID: projectID}).Preload("Project").Preload("Role").First(&projectMember).Error
	return &projectMember, errors.Is(err, gorm.ErrRecordNotFound), err
}

func (projectMemberDao) DeleteProjectMember(projectMember *models.ProjectMember) error {
	result := db.GetDB().Where(models.ProjectMember{UserID: projectMember.UserID, ProjectID: projectMember.ProjectID}).Delete(&models.ProjectMember{})
	return result.Error
}
