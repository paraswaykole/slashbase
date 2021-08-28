package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/models"
)

var userDao daos.UserDao
var projectDao daos.ProjectDao

const (
	USER_SESSION = "USER_SESSION"
)

// FindUserMiddleware is find authenticated user before sending the request to next handler
func FindUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		auth := c.GetHeader("Authorization")
		if auth != "" && strings.HasPrefix(auth, "Bearer ") {
			tokenString = strings.ReplaceAll(auth, "Bearer ", "")
		} else {
			tokenString, _ = c.Cookie("token")
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

func GetAuthUserHasRolesForProject(c *gin.Context, projectID string, hasRoles []string) (bool, error) {
	authUser := GetAuthUser(c)
	projectMembers, err := projectDao.GetProjectMembersForUser(authUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   "There was some problem",
		})
		return false, err
	}
	for _, pMember := range *projectMembers {
		if pMember.ProjectID == projectID {
			if utils.ExistsIn(pMember.Role, &hasRoles) {
				return true, nil
			}
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   "not allowed",
			})
			return false, nil
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"error":   "not allowed",
	})
	return false, nil
}
