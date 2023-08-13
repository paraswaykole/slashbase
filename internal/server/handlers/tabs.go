package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/views"
	"github.com/slashbaseide/slashbase/internal/server/controllers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
)

type TabsHandlers struct{}

var tabController controllers.TabsController

func (TabsHandlers) CreateNewTab(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var createBody struct {
		DBConnectionId string `json:"dbConnectionId"`
		TabType        string `json:"tabType"`
		Modelschema    string `json:"modelschema"`
		Modelname      string `json:"modelname"`
		QueryID        string `json:"queryId"`
		Query          string `json:"query"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return fiber.ErrBadRequest
	}
	tab, err := tabController.CreateTab(authUser.ID, createBody.DBConnectionId, createBody.TabType, createBody.Modelschema, createBody.Modelname, createBody.QueryID, createBody.Query)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildTabView(&tab.Tab),
	})
}

func (TabsHandlers) GetTabsByDBConnection(c *fiber.Ctx) error {
	dbConnectionId := c.Params("dbConnId")
	authUser := middlewares.GetAuthUser(c)
	tabs, err := tabController.GetTabsByDBConnection(authUser.ID, dbConnectionId)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	tabViews := []views.TabView{}
	for _, t := range *tabs {
		tabViews = append(tabViews, *views.BuildTabView(&t.Tab))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    tabViews,
	})
}

func (TabsHandlers) UpdateTab(c *fiber.Ctx) error {
	authUser := middlewares.GetAuthUser(c)
	var updateBody struct {
		DBConnectionID string                 `json:"dbConnectionId"`
		TabID          string                 `json:"tabId"`
		TabType        string                 `json:"tabType"`
		Metadata       map[string]interface{} `json:"metadata"`
	}
	if err := c.BodyParser(&updateBody); err != nil {
		return fiber.ErrBadRequest
	}
	tab, err := tabController.UpdateTab(authUser.ID, updateBody.DBConnectionID, updateBody.TabID, updateBody.TabType, updateBody.Metadata)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildTabView(&tab.Tab),
	})
}

func (TabsHandlers) CloseTab(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	tabID := c.Params("tabId")
	authUser := middlewares.GetAuthUser(c)
	err := tabController.CloseTab(authUser.ID, dbConnID, tabID)
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
