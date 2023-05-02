package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/server/models"
)

type RoleView struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Permissions []RolePermissionView `json:"permissions,omitempty"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

type RolePermissionView struct {
	ID        string    `json:"id"`
	RoleID    string    `json:"roleId"`
	Name      string    `json:"name"`
	Value     bool      `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildRole(role *models.Role) RoleView {
	roleView := RoleView{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
	if len(role.Permissions) > 0 {
		pViews := []RolePermissionView{}
		for _, rp := range role.Permissions {
			pViews = append(pViews, BuildRolePermission(&rp))
		}
		roleView.Permissions = pViews
	}
	return roleView
}

func BuildRolePermission(rp *models.RolePermission) RolePermissionView {
	rpView := RolePermissionView{
		ID:        rp.ID,
		RoleID:    rp.RoleID,
		Name:      rp.Name,
		Value:     rp.Value,
		CreatedAt: rp.CreatedAt,
		UpdatedAt: rp.UpdatedAt,
	}
	return rpView
}
