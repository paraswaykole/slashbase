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

func (d RoleDao) GetAllRoles() (*[]models.Role, error) {
	var roles []models.Role
	err := db.GetDB().Model(models.Role{}).Find(&roles).Error
	return &roles, err
}

func (d RoleDao) DeleteRoleById(roleId string) error {
	err := db.GetDB().Where(models.Role{ID: roleId}).Delete(&models.Role{}).Error
	return err
}
