package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type RoleController struct{}

func (RoleController) GetAllRoles(user *models.User) (*[]models.Role, *[]models.RolePermission, error) {

	if !user.IsRoot {
		return nil, nil, errors.New("not allowed")
	}

	roles, err := dao.Role.GetAllRoles()
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	rolePermissions, err := dao.RolePermission.GetAllRolePermissions()
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	return roles, rolePermissions, nil
}

func (RoleController) AddRole(user *models.User, name string) (*models.Role, error) {

	if !user.IsRoot {
		return nil, errors.New("not allowed")
	}

	role := models.NewRole(name)
	err := dao.Role.CreateRole(role)
	if err != nil {
		return nil, errors.New("cannot create role: " + name)
	}
	return role, nil
}

func (RoleController) DeleteRole(user *models.User, roleID string) error {

	if !user.IsRoot {
		return errors.New("not allowed")
	}

	role, err := dao.Role.GetAdminRole()
	if err != nil {
		return errors.New("there was some problem")
	}

	if role.ID == roleID {
		return errors.New("cannot delete admin role")
	}

	err = dao.Role.DeleteRoleByID(roleID)
	if err != nil {
		return errors.New("cannot delete role")
	}

	return nil
}

func (RoleController) AddOrUpdateRolePermission(user *models.User, roleID, name string, value bool) (*models.RolePermission, error) {

	if !user.IsRoot {
		return nil, errors.New("not allowed")
	}

	rp, err := dao.RolePermission.UpdateRolePermission(roleID, name, value)
	if err != nil {
		return nil, errors.New("cannot update role permission: " + name)
	}
	return rp, nil
}
