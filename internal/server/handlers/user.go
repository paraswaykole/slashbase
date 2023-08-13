package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
	"github.com/slashbaseide/slashbase/internal/server/views"
)

type UserHandlers struct{}

var userController controllers.UserController

func (UserHandlers) LoginUser(c *fiber.Ctx) error {
	var loginBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginBody); err != nil {
		return fiber.ErrBadRequest
	}

	userSession, err := userController.LoginUser(loginBody.Email, loginBody.Password)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     config.SESSION_COOKIE_NAME,
		Value:    userSession.GetAuthToken(),
		MaxAge:   config.SESSION_COOKIE_MAX_AGE,
		Path:     "/",
		Domain:   utils.ExtractDomainFromHost(string(c.Request().Host())),
		Secure:   c.Context().IsTLS(),
		HTTPOnly: true,
		SameSite: "strict",
	})
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildUserSession(userSession),
	})
}

func (UserHandlers) CheckAuth(c *fiber.Ctx) error {
	tokenString := c.Cookies(config.SESSION_COOKIE_NAME)
	if tokenString != "" {
		return c.JSON(map[string]interface{}{
			"success": true,
		})
	}
	return c.JSON(map[string]interface{}{
		"success": false,
	})
}

func (UserHandlers) EditAccount(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var userBody struct {
		Name            string `json:"name"`
		ProfileImageURL string `json:"profileImageUrl"`
	}
	if err := c.BodyParser(&userBody); err != nil {
		return fiber.ErrBadRequest
	}
	err := userController.EditAccount(authUser, userBody.Name, userBody.ProfileImageURL)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildUser(authUser),
	})
}

func (UserHandlers) ChangePassword(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var body struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	err := userController.ChangePassword(authUser, body.OldPassword, body.NewPassword)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}

func (UserHandlers) GetUsers(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   "invalid offset",
		})
	}

	searchTerm := c.Query("search", "")

	users, next, err := userController.GetUsersPaginated(authUser, searchTerm, offset)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	userViews := []views.UserView{}
	for _, user := range *users {
		userViews = append(userViews, views.BuildUser(&user))
	}

	return c.JSON(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"list": userViews,
			"next": next,
		},
	})
}

func (UserHandlers) AddUsers(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var addUserBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&addUserBody); err != nil {
		return fiber.ErrBadRequest
	}
	err := userController.AddUser(authUser, addUserBody.Email, addUserBody.Password)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}

func (UserHandlers) Logout(c *fiber.Ctx) error {
	authUserSession := middlewares.GetAuthSession(c)
	authUserSession.SetInactive()
	c.Cookie(&fiber.Cookie{
		Name:     config.SESSION_COOKIE_NAME,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   utils.ExtractDomainFromHost(string(c.Request().Host())),
		Secure:   c.Context().IsTLS(),
		HTTPOnly: true,
		SameSite: "strict",
	})
	return c.JSON(map[string]interface{}{
		"success": true,
	})
}
