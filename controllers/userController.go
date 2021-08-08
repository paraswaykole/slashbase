package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type UserController struct{}

var userDao daos.UserDao

func (uc UserController) LoginUser(c *gin.Context) {
	var loginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&loginBody)
	usr, err := userDao.GetUserByEmail(loginBody.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   "Invalid User",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	if usr.VerifyPassword(loginBody.Password) {
		userSession, _ := models.NewUserSession(usr.ID)
		err = userDao.CreateUserSession(userSession)
		userSession.User = *usr
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   "There was some problem",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    views.BuildUserSession(userSession),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"error":   "Invalid Login",
	})
	return
}

func (uc UserController) AddUser(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var addUserBody struct {
		Email string `json:"email"`
	}
	c.BindJSON(&addUserBody)
	if !authUser.RootUser {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "Not Allowed.",
		})
	}
	usr, err := userDao.GetUserByEmail(addUserBody.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usr, err = models.NewUser(addUserBody.Email, utils.RandStringUnsafe(10))
			if err == nil {
				err = userDao.CreateUser(usr)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"success": false,
						"error":   "There was some problem",
					})
					return
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"error":   err,
				})
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   "There was some problem",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (uc UserController) Logout(c *gin.Context) {
	authUserSession := middlewares.GetAuthSession(c)
	authUserSession.SetInActive()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}
