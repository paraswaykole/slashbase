package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:     "Slashbase Server",
		ReadTimeout: time.Second * time.Duration(60),
	})
	return app
}
