package dao

import (
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"gorm.io/gorm/clause"
)

type rolePermissionDao struct{}

var RolePermission rolePermissionDao

func (rolePermissionDao) CreateRolePermission(rp *models.RolePermission) error {
	err := db.GetDB().Create(rp).Error
	return err
}

func (rolePermissionDao) GetRolePermissionsForRole(roleID string) (*[]models.RolePermission, error) {
	var rps []models.RolePermission
	err := db.GetDB().Where(models.RolePermission{RoleID: roleID}).Find(&rps).Error
	return &rps, err
}

func (rolePermissionDao) GetAllRolePermissions() (*[]models.RolePermission, error) {
	var rps []models.RolePermission
	err := db.GetDB().Model(models.RolePermission{}).Find(&rps).Error
	return &rps, err
}

func (rolePermissionDao) UpdateRolePermission(roleID, name string, value bool) (*models.RolePermission, error) {
	rolePermission := models.NewRolePermission(roleID, name, value)
	err := db.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "role_id"}, {Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(rolePermission).Error
	return rolePermission, err
}
