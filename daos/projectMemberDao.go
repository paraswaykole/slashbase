package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

func (d ProjectDao) CreateProjectMembers(projectMembers *[]models.ProjectMember) error {
	result := db.GetDB().Create(projectMembers)
	return result.Error
}

func (d ProjectDao) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{ProjectID: projectID}).Preload("User").Find(&projectMembers).Error
	return &projectMembers, err
}

func (d ProjectDao) GetProjectMembersForUser(userID string) (*[]models.ProjectMember, error) {
	var projectMembers []models.ProjectMember
	err := db.GetDB().Where(models.ProjectMember{UserID: userID}).Preload("Project").Find(&projectMembers).Error
	return &projectMembers, err
}
