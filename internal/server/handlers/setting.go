package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
)

type SettingHandlers struct{}

var settingController controllers.SettingController

func (SettingHandlers) GetSingleSetting(c *fiber.Ctx) error {
	name := c.Query("name")
	value, err := settingController.GetSingleSetting(name)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    value,
	})
}

func (SettingHandlers) UpdateSingleSetting(c *fiber.Ctx) error {
	var reqBody struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return fiber.ErrBadRequest
	}
	err := settingController.UpdateSingleSetting(reqBody.Name, reqBody.Value)
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
