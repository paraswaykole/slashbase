package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
)

type ConsoleHandlers struct{}

var consoleController controllers.ConsoleController

func (ConsoleHandlers) RunCommandEvent(c *fiber.Ctx) error {
	var body struct {
		DBConnectionID string `json:"dbConnectionId"`
		CmdString      string `json:"cmd"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	analytics.SendRunCommandEvent()
	output := consoleController.RunCommand(body.DBConnectionID, body.CmdString)
	return c.JSON(map[string]interface{}{
		"success": false,
		"data":    output,
	})
}
