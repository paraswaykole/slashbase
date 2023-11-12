package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
)

type AIHandlers struct{}

var aiController controllers.AIController

func (AIHandlers) GenerateSQL(c *fiber.Ctx) error {
	var body struct {
		DBConnectionID string `json:"dbConnectionId"`
		Text           string `json:"text"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	analytics.SendAISQLGeneratedEvent()
	output, err := aiController.GenerateSQL(body.DBConnectionID, body.Text)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    output,
	})
}

func (AIHandlers) ListSupportedAIModels(c *fiber.Ctx) error {
	output := aiController.GetModels()
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    output,
	})
}
