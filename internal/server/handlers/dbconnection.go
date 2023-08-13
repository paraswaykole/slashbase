package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slashbaseide/slashbase/internal/common/controllers"
	"github.com/slashbaseide/slashbase/internal/common/views"
)

type DBConnectionHandlers struct{}

var dbConnController controllers.DBConnectionController

func (DBConnectionHandlers) CreateDBConnection(c *fiber.Ctx) error {

	var createBody struct {
		ProjectID   string `json:"projectId"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Scheme      string `json:"scheme"`
		Host        string `json:"host"`
		Port        string `json:"port"`
		Password    string `json:"password"`
		User        string `json:"user"`
		DBName      string `json:"dbname"`
		UseSSL      bool   `json:"useSSL"`
		UseSSH      string `json:"useSSH"`
		SSHHost     string `json:"sshHost"`
		SSHUser     string `json:"sshUser"`
		SSHPassword string `json:"sshPassword"`
		SSHKeyFile  string `json:"sshKeyFile"`
		IsTest      bool   `json:"isTest"`
	}
	if err := c.BodyParser(&createBody); err != nil {
		return fiber.ErrBadRequest
	}
	dbConn, err := dbConnController.CreateDBConnection(
		createBody.ProjectID,
		createBody.Name,
		createBody.Type,
		createBody.Scheme,
		createBody.Host,
		createBody.Port,
		createBody.User,
		createBody.Password,
		createBody.DBName,
		createBody.UseSSH,
		createBody.SSHHost,
		createBody.SSHUser,
		createBody.SSHPassword,
		createBody.SSHKeyFile,
		createBody.UseSSL,
		createBody.IsTest,
	)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	if dbConn == nil {
		return c.JSON(map[string]interface{}{
			"success": true,
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildDBConnection(dbConn),
	})
}

func (DBConnectionHandlers) GetDBConnections(c *fiber.Ctx) error {
	dbConns, err := dbConnController.GetDBConnections()
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	dbConnViews := []views.DBConnectionView{}
	for _, dbConn := range dbConns {
		dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    dbConnViews,
	})
}

func (DBConnectionHandlers) DeleteDBConnection(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	err := dbConnController.DeleteDBConnection(dbConnID)
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

func (DBConnectionHandlers) GetSingleDBConnection(c *fiber.Ctx) error {
	dbConnID := c.Params("dbConnId")
	dbConn, err := dbConnController.GetSingleDBConnection(dbConnID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    views.BuildDBConnection(dbConn),
	})
}

func (DBConnectionHandlers) GetDBConnectionsByProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	dbConns, err := dbConnController.GetDBConnectionsByProject(projectID)
	if err != nil {
		return c.JSON(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}
	dbConnViews := []views.DBConnectionView{}
	for _, dbConn := range dbConns {
		dbConnViews = append(dbConnViews, views.BuildDBConnection(dbConn))
	}
	return c.JSON(map[string]interface{}{
		"success": true,
		"data":    dbConnViews,
	})
}

func (DBConnectionHandlers) CheckDBConnection(c *fiber.Ctx) error {
	dbConnectionID := c.Params("dbConnId")
	err := dbConnController.CheckDBConnection(dbConnectionID)
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
