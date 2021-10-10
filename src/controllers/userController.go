package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/views"
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
}

func (uc UserController) EditAccount(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var userBody struct {
		Name            string `json:"name"`
		ProfileImageURL string `json:"profileImageUrl"`
	}
	c.BindJSON(&userBody)

	err := userDao.EditUser(authUser.ID, userBody.Name, userBody.ProfileImageURL)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	authUser.FullName = sql.NullString{
		String: userBody.Name,
		Valid:  userBody.Name != "",
	}
	authUser.ProfileImageURL = sql.NullString{
		String: userBody.ProfileImageURL,
		Valid:  userBody.ProfileImageURL != "",
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildUser(authUser),
	})
}

func (uc UserController) GetUsers(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	if !authUser.IsRoot {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "not allowed",
		})
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "invalid offset",
		})
	}

	users, err := userDao.GetUsersPaginated(offset)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
	}

	userViews := []views.UserView{}
	next := -1
	for i, user := range *users {
		userViews = append(userViews, views.BuildUser(&user))
		if i == config.PAGINATION_COUNT-1 {
			next = next + config.PAGINATION_COUNT
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"list": userViews,
			"next": next,
		},
	})

}

func (uc UserController) AddUser(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	var addUserBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&addUserBody)
	if !authUser.IsRoot {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "not allowed",
		})
	}
	usr, err := userDao.GetUserByEmail(addUserBody.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usr, err = models.NewUser(addUserBody.Email, addUserBody.Password)
			if err != nil {
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
	err = userDao.CreateUser(usr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (uc UserController) Logout(c *gin.Context) {
	authUserSession := middlewares.GetAuthSession(c)
	authUserSession.SetInActive()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
