package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/controllers"
	"slashbase.com/backend/internal/middlewares"
	"slashbase.com/backend/internal/views"
)

type RoleHandlers struct{}

var roleController controllers.RoleController

func (RoleHandlers) GetAllRoles(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	allRoles, allRolePermissions, err := roleController.GetAllRoles(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
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
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    roleViews,
	})
}

func (RoleHandlers) AddRole(c *gin.Context) {
	var reqBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&reqBody)
	authUser := middlewares.GetAuthUser(c)
	role, err := roleController.AddRole(authUser, reqBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildRole(role),
	})
}

func (RoleHandlers) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	authUser := middlewares.GetAuthUser(c)
	err := roleController.DeleteRole(authUser, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (RoleHandlers) UpdateRolePermission(c *gin.Context) {
	roleID := c.Param("id")
	var reqBody struct {
		Name  string `json:"name"`
		Value bool   `json:"value"`
	}
	c.BindJSON(&reqBody)
	authUser := middlewares.GetAuthUser(c)
	rp, err := roleController.AddOrUpdateRolePermission(authUser, roleID, reqBody.Name, reqBody.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildRolePermission(rp),
	})
}
