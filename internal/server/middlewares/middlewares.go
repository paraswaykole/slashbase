package middlewares

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

const (
	USER_SESSION = "USER_SESSION"
)

func APIResponseMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(code).JSON(map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			})
		}
		if err == nil {
			response := c.Response()
			body := response.Body()
			var data interface{}
			json.Unmarshal(body, &data)

			if resMap, ok := data.(map[string]interface{}); ok {
				if _, ok := resMap["success"]; ok {
					return c.JSON(resMap)
				}
			}

			return c.JSON(map[string]interface{}{
				"success": true,
				"data":    data,
			})
		}
		return nil
	}
}

// FindUserMiddleware is to find authenticated user before sending the request to next handler
func FindUserMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies(config.SESSION_COOKIE_NAME, "")
		if tokenString == "" {
			auth := c.GetReqHeaders()["Authorization"]
			if auth != "" && strings.HasPrefix(auth, "Bearer ") {
				tokenString = strings.ReplaceAll(auth, "Bearer ", "")
			}
		}
		if tokenString != "" {
			userSession, err := dao.User.GetUserSessionFromAuthToken(tokenString)
			if err != nil {
				return c.Next()
			}
			c.Context().SetUserValue(USER_SESSION, userSession)
			return c.Next()
		}
		return c.Next()
	}
}

// AuthUserMiddleware is to check if authUser is present else returns unauthorized error
func AuthUserMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if value := c.Context().UserValue(USER_SESSION); value != nil {
			return c.Next()
		}
		return fiber.ErrUnauthorized
	}
}

func GetAuthSession(c *fiber.Ctx) *models.UserSession {
	if session := c.Context().UserValue(USER_SESSION); session.(*models.UserSession) != nil {
		return session.(*models.UserSession)
	}
	return nil
}

func GetAuthUser(c *fiber.Ctx) *models.User {
	if session := c.Context().UserValue(USER_SESSION); session != nil {
		authUserSession := session.(*models.UserSession)
		return &authUserSession.User
	}
	return nil
}

func GetAuthUserProjectIds(c *fiber.Ctx) *[]string {
	authUserSession := c.Context().UserValue(USER_SESSION).(*models.UserSession)
	projectIDs := []string{}
	for _, project := range authUserSession.User.Projects {
		projectIDs = append(projectIDs, project.ID)
	}
	return &projectIDs
}
