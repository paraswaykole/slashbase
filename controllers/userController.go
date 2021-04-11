package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models/user"
	"slashbase.com/backend/views"
)

// UserController is parent for all methods below
type UserController struct{}

var userDao user.Dao

func (uc UserController) LoginUser(c *gin.Context) {
	var loginCmd struct {
		Email string `json:"email"`
	}
	c.BindJSON(&loginCmd)
	usr, err := userDao.GetUserByEmail(loginCmd.Email)
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
	userSession, _ := user.NewUserSession(usr.ID)
	err = userDao.CreateUserSession(userSession)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	// TODO: send email
	fmt.Println("Magic Link: " + userSession.GetMagicLink())
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (uc UserController) RegisterUser(c *gin.Context) {
	var registerCmd struct {
		Email string `json:"email"`
	}
	c.BindJSON(&registerCmd)
	usr, err := userDao.GetUserByEmail(registerCmd.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usr, err = user.NewUser(registerCmd.Email)
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
	userSession, _ := user.NewUserSession(usr.ID)
	err = userDao.CreateUserSession(userSession)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	// TODO: send email
	fmt.Println("Magic Link: " + userSession.GetMagicLink())
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
	return
}

func (uc UserController) VerifySession(c *gin.Context) {
	var verifyCmd struct {
		Token string `json:"token"`
	}
	c.BindJSON(&verifyCmd)
	userSession, err := userDao.GetUserSessionFromMagicLinkToken(verifyCmd.Token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	userSession.SetActive()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildUserSession(userSession),
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
