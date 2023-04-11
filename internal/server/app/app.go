package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:     "Slashbase Server",
		ReadTimeout: time.Second * time.Duration(60),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return app
}
