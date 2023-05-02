package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/slashbaseide/slashbase/internal/common/config"
)

func CreateFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:     "Slashbase Server",
		ReadTimeout: time.Second * time.Duration(60),
	})

	corsConfig := cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}
	if !config.IsLive() {
		corsConfig.AllowOrigins = "http://localhost:5173"
	}

	app.Use(cors.New(corsConfig))

	return app
}
