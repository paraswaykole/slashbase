package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
	"github.com/slashbaseide/slashbase/internal/server/views"
)

type RoleHandlers struct{}

var roleController controllers.RoleController

func (RoleHandlers) GetAllRoles(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	allRoles, allRolePermissions, err := roleController.GetAllRoles(authUser)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	rolePermissionViews := map[string][]views.RolePermissionView{}

	for _, rp := range *allRolePermissions {
		rpview := views.BuildRolePermission(&rp)
		if _, exists := rolePermissionViews[rp.RoleID]; !exists {
			rolePermissionViews[rpview.RoleID] = []views.RolePermissionView{}
		}
		rolePermissionViews[rpview.RoleID] = append(rolePermissionViews[rpview.RoleID], rpview)
	}

	roleViews := []views.RoleView{}
	for _, r := range *allRoles {
		roleView := views.BuildRole(&r)
		if rpviews, exists := rolePermissionViews[r.ID]; exists {
			roleView.Permissions = rpviews
		}
		roleViews = append(roleViews, roleView)
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    roleViews,
	})
}

func (RoleHandlers) AddRole(c *fiber.Ctx) error {
	var reqBody struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	authUser := middlewares.GetAuthUser(c)
	role, err := roleController.AddRole(authUser, reqBody.Name)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildRole(role),
	})
}

func (RoleHandlers) DeleteRole(c *fiber.Ctx) error {
	roleID := c.Params("id")
	authUser := middlewares.GetAuthUser(c)
	err := roleController.DeleteRole(authUser, roleID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}

func (RoleHandlers) UpdateRolePermission(c *fiber.Ctx) error {
	roleID := c.Params("id")
	var reqBody struct {
		Name  string `json:"name"`
		Value bool   `json:"value"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	authUser := middlewares.GetAuthUser(c)
	rp, err := roleController.AddOrUpdateRolePermission(authUser, roleID, reqBody.Name, reqBody.Value)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildRolePermission(rp),
	})
}
