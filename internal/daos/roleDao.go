package daos

import (
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

type RoleDao struct{}

func (d RoleDao) CreateRole(role *models.Role) error {
	err := db.GetDB().Create(role).Error
	return err
}

func (d RoleDao) GetAdminRole() (*models.Role, error) {
	var role models.Role
	err := db.GetDB().Where(models.Role{Name: models.ROLE_ADMIN}).FirstOrCreate(&role).Error
	return &role, err
}
