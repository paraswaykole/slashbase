package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
	"github.com/slashbaseide/slashbase/internal/common/views"
)

type TabsHandlers struct{}

var tabController controllers.TabsController

func (TabsHandlers) CreateNewTab(c *fiber.Ctx) error {
	var createBody struct {
		DBConnectionId string `json:"dbConnectionId"`
		TabType        string `json:"tabType"`
		Modelschema    string `json:"modelschema"`
		Modelname      string `json:"modelname"`
		QueryID        string `json:"queryId"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	tab, err := tabController.CreateTab(createBody.DBConnectionId, createBody.TabType, createBody.Modelschema, createBody.Modelname, createBody.QueryID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildTabView(tab),
	})
}

func (TabsHandlers) GetTabsByDBConnection(c *fiber.Ctx) error {
	dbConnectionId := c.Params("dbConnId")
	tabs, err := tabController.GetTabsByDBConnection(dbConnectionId)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	tabViews := []views.TabView{}
	for _, t := range *tabs {
		tabViews = append(tabViews, *views.BuildTabView(&t))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    tabViews,
	})
}

func (TabsHandlers) UpdateTab(c *fiber.Ctx) error {
	var updateBody struct {
		DBConnectionID string                 `json:"dbConnectionId"`
		TabID          string                 `json:"tabId"`
		TabType        string                 `json:"tabType"`
		Metadata       map[string]interface{} `json:"metadata"`
	}
	if err := c.BodyParser(&updateBody); err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	tab, err := tabController.UpdateTab(updateBody.DBConnectionID, updateBody.TabID, updateBody.TabType, updateBody.Metadata)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildTabView(tab),
	})
}

func (TabsHandlers) CloseTab(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	tabID := c.Params("tabId")
	err := tabController.CloseTab(dbConnID, tabID)
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
