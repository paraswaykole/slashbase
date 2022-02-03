package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/controllers"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/views"
)

type UserRoutes struct{}

var userController controllers.UserController

func (ur UserRoutes) LoginUser(c *gin.Context) {
	var loginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&loginBody)
	userSession, err := userController.LoginUser(loginBody.Email, loginBody.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildUserSession(userSession),
	})
}

func (ur UserRoutes) EditAccount(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var userBody struct {
		Name            string `json:"name"`
		ProfileImageURL string `json:"profileImageUrl"`
	}
	c.BindJSON(&userBody)

	err := userController.EditAccount(authUser, userBody.Name, userBody.ProfileImageURL)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildUser(authUser),
	})
}

func (ur UserRoutes) ChangePassword(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var body struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	c.BindJSON(&body)

	err := userController.ChangePassword(authUser, body.OldPassword, body.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ur UserRoutes) GetUsers(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "invalid offset",
		})
	}

	searchTerm := c.DefaultQuery("search", "")

	users, next, err := userController.GetUsersPaginated(authUser, searchTerm, offset)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	userViews := []views.UserView{}
	for _, user := range *users {
		userViews = append(userViews, views.BuildUser(&user))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"list": userViews,
			"next": next,
		},
	})

}

func (ur UserRoutes) AddUser(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var addUserBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&addUserBody)
	err := userController.AddUser(authUser, addUserBody.Email, addUserBody.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ur UserRoutes) Logout(c *gin.Context) {
	authUserSession := middlewares.GetAuthSession(c)
	authUserSession.SetInActive()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
