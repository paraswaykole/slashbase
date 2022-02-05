package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/models"
)

var userDao daos.UserDao

const (
	USER_SESSION = "USER_SESSION"
)

// FindUserMiddleware is find authenticated user before sending the request to next handler
func FindUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie(config.SESSION_COOKIE_NAME)
		if err == http.ErrNoCookie {
			auth := c.GetHeader("Authorization")
			if auth != "" && strings.HasPrefix(auth, "Bearer ") {
				tokenString = strings.ReplaceAll(auth, "Bearer ", "")
			}
		}
		if tokenString != "" {
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
	}
}

// AuthUserMiddleware is checks if authUser is present else returns unauthorized error
func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(USER_SESSION); exists {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
		})
		c.Abort()
	}
}

func GetAuthSession(c *gin.Context) *models.UserSession {
	authUserSession := c.MustGet(USER_SESSION).(*models.UserSession)
	return authUserSession
}

func GetAuthUser(c *gin.Context) *models.User {
	authUserSession := c.MustGet(USER_SESSION).(*models.UserSession)
	return &authUserSession.User
}

func GetAuthUserProjectIds(c *gin.Context) *[]string {
	authUserSession := c.MustGet(USER_SESSION).(*models.UserSession)
	projectIDs := []string{}
	for _, project := range authUserSession.User.Projects {
		projectIDs = append(projectIDs, project.ID)
	}
	return &projectIDs
}
