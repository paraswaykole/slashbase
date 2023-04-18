package dao

import (
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type roleDao struct{}

var Role roleDao

func (roleDao) CreateRole(role *models.Role) error {
	err := db.GetDB().Create(role).Error
	return err
}

func (roleDao) GetAdminRole() (*models.Role, error) {
	var role models.Role
	err := db.GetDB().Where(models.Role{Name: models.ROLE_ADMIN}).Attrs(models.NewRole(models.ROLE_ADMIN)).FirstOrCreate(&role).Error
	return &role, err
}

func (roleDao) GetRoleByID(roleID string) (*models.Role, error) {
	var role models.Role
	err := db.GetDB().Where(models.Role{ID: roleID}).First(&role).Error
	return &role, err
}

func (roleDao) GetAllRoles() (*[]models.Role, error) {
	var roles []models.Role
	err := db.GetDB().Find(&roles).Error
	return &roles, err
}

func (roleDao) DeleteRoleByID(roleId string) error {
	err := db.GetDB().Where(models.Role{ID: roleId}).Delete(&models.Role{}).Error
	return err
}
