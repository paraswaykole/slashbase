package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/middlewares"
	"slashbase.com/backend/internal/views"
)

type RoleRoutes struct{}

func (rr RoleRoutes) GetAllRoles(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	allRoles, err := projectController.GetAllRoles(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	roleViews := []views.RoleView{}
	for _, r := range *allRoles {
		roleViews = append(roleViews, views.BuildRole(&r))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    roleViews,
	})
}

func (rr RoleRoutes) AddRole(c *gin.Context) {
	var reqBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&reqBody)
	authUser := middlewares.GetAuthUser(c)
	role, err := projectController.AddRole(authUser, reqBody.Name)
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

func (rr RoleRoutes) DeleteRole(c *gin.Context) {
	roleID := c.Param("id")
	authUser := middlewares.GetAuthUser(c)
	err := projectController.DeleteRole(authUser, roleID)
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
