package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/controllers"
	"slashbase.com/backend/internal/middlewares"
	"slashbase.com/backend/internal/utils"
	"slashbase.com/backend/internal/views"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(config.SESSION_COOKIE_NAME, userSession.GetAuthToken(), config.SESSION_COOKIE_MAX_AGE, "/", utils.GetRequestCookieHost(c.Request), utils.GetRequestScheme(c.Request) == "https", true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildUserSession(userSession),
	})
}

func (ur UserRoutes) CheckAuth(c *gin.Context) {
	tokenString, _ := c.Cookie("session")
	if tokenString != "" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
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
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (ur UserRoutes) GetUsers(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid offset",
		})
	}

	searchTerm := c.DefaultQuery("search", "")

	users, next, err := userController.GetUsersPaginated(authUser, searchTerm, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (ur UserRoutes) Logout(c *gin.Context) {
	authUserSession := middlewares.GetAuthSession(c)
	authUserSession.SetInActive()
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(config.SESSION_COOKIE_NAME, "", -1, "/", utils.GetRequestCookieHost(c.Request), utils.GetRequestScheme(c.Request) == "https", true)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
