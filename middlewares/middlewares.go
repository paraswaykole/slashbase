package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/models/user"
)

var userDao user.Dao

const (
	USER_SESSION = "USER_SESSION"
)

// FindUserMiddleware is find authenticated user before sending the request to next handler
func FindUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth != "" && strings.HasPrefix(auth, "Bearer ") {
			tokenString := strings.ReplaceAll(auth, "Bearer ", "")
			userSession, err := userDao.GetUserSessionFromAuthToken(tokenString)
			if err != nil {
				c.Next()
				return
			}
			c.Set(USER_SESSION, userSession)
			c.Next()
			return
		}
		c.Next()
		return
	}
}

// AuthUserMiddleware is checks if authUser is present else returns unauthorized error
func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(USER_SESSION); exists {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}

func GetAuthSession(c *gin.Context) *user.UserSession {
	authUserSession := c.MustGet(USER_SESSION).(*user.UserSession)
	return authUserSession
}

func GetAuthUser(c *gin.Context) *user.User {
	authUserSession := c.MustGet(USER_SESSION).(*user.UserSession)
	return &authUserSession.User
}
