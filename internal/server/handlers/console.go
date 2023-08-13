package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
)

type ConsoleHandlers struct{}

var consoleController controllers.ConsoleController

func (ConsoleHandlers) RunCommand(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var body struct {
		DBConnectionID string `json:"dbConnectionId"`
		CmdString      string `json:"cmd"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	analytics.SendRunCommandEvent()
	output := consoleController.RunCommand(authUser, body.DBConnectionID, body.CmdString)
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    output,
	})
}
