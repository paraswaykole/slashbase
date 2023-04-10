package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/server/handlers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	{
		api.Get("health", healthCheck)
		userGroup := api.Group("user")
		{
			userHandlers := new(handlers.UserHandlers)
			userGroup.Post("/login", userHandlers.LoginUser)
			userGroup.Get("/checkauth", userHandlers.CheckAuth)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.Post("/edit", userHandlers.EditAccount)
			userGroup.Post("/password", userHandlers.ChangePassword)
			userGroup.Post("/add", userHandlers.AddUsers)
			userGroup.Get("/all", userHandlers.GetUsers)
			userGroup.Get("/logout", userHandlers.Logout)
		}
	}
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"success": true,
		"version": config.GetConfig().Version,
	})
}
