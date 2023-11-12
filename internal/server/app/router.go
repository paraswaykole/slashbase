package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/server/handlers"
	"github.com/slashbaseide/slashbase/internal/server/middlewares"
)

func SetupRoutes(app *fiber.App, assets embed.FS) {
	api := app.Group("/api/v1")
	{
		api.Use(middlewares.APIResponseMiddleware())
		api.Get("health", healthCheck)
		userGroup := api.Group("user")
		{
			userHandlers := new(handlers.UserHandlers)
			userGroup.Post("/login", userHandlers.LoginUser)
			userGroup.Get("/checkauth", userHandlers.CheckAuth)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.Post("/edit", userHandlers.EditAccount)
			userGroup.Post("/password", userHandlers.ChangePassword)
			userGroup.Post("/add", userHandlers.AddUsers)
			userGroup.Get("/all", userHandlers.GetUsers)
			userGroup.Get("/logout", userHandlers.Logout)
		}
		projectGroup := api.Group("project")
		{
			projectHandlers := new(handlers.ProjectHandlers)
			projectGroup.Use(middlewares.FindUserMiddleware())
			projectGroup.Use(middlewares.AuthUserMiddleware())
			projectGroup.Post("/create", projectHandlers.CreateProject)
			projectGroup.Get("/all", projectHandlers.GetProjects)
			projectGroup.Delete("/:projectId", projectHandlers.DeleteProject)
			projectGroup.Post("/:projectId/members/create", projectHandlers.AddProjectMember)
			projectGroup.Get("/:projectId/members", projectHandlers.GetProjectMembers)
			projectGroup.Delete("/:projectId/members/:userId", projectHandlers.DeleteProjectMember)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnectionHandler := new(handlers.DBConnectionHandlers)
			dbConnGroup.Use(middlewares.FindUserMiddleware())
			dbConnGroup.Use(middlewares.AuthUserMiddleware())
			dbConnGroup.Post("/create", dbConnectionHandler.CreateDBConnection)
			dbConnGroup.Get("/all", dbConnectionHandler.GetDBConnections)
			dbConnGroup.Get("/project/:projectId", dbConnectionHandler.GetDBConnectionsByProject)
			dbConnGroup.Get("/:dbConnId", dbConnectionHandler.GetSingleDBConnection)
			dbConnGroup.Get("/check/:dbConnId", dbConnectionHandler.CheckDBConnection)
			dbConnGroup.Delete("/:dbConnId", dbConnectionHandler.DeleteDBConnection)
		}
		queryGroup := api.Group("query")
		{
			queryHandlers := new(handlers.QueryHandlers)
			queryGroup.Use(middlewares.FindUserMiddleware())
			queryGroup.Use(middlewares.AuthUserMiddleware())
			queryGroup.Post("/run", queryHandlers.RunQuery)
			queryGroup.Post("/save/:dbConnId", queryHandlers.SaveDBQuery)
			queryGroup.Get("/getall/:dbConnId", queryHandlers.GetDBQueriesInDBConnection)
			queryGroup.Get("/get/:queryId", queryHandlers.GetSingleDBQuery)
			queryGroup.Delete("/delete/:queryId", queryHandlers.DeleteDBQuery)
			queryGroup.Get("/history/:dbConnId", queryHandlers.GetQueryHistoryInDBConnection)
			dataGroup := queryGroup.Group("data")
			{
				dataGroup.Get("/:dbConnId", queryHandlers.GetData)
				dataGroup.Post("/:dbConnId/single", queryHandlers.UpdateSingleData)
				dataGroup.Post("/:dbConnId/add", queryHandlers.AddData)
				dataGroup.Post("/:dbConnId/delete", queryHandlers.DeleteData)
			}
			dataModelGroup := queryGroup.Group("datamodel")
			{
				dataModelGroup.Get("/all/:dbConnId", queryHandlers.GetDataModels)
				dataModelGroup.Get("/single/:dbConnId", queryHandlers.GetSingleDataModel)
				dataModelGroup.Post("/single/addfield", queryHandlers.AddSingleDataModelField)
				dataModelGroup.Post("/single/deletefield", queryHandlers.DeleteSingleDataModelField)
				dataModelGroup.Post("/single/addindex", queryHandlers.AddSingleDataModelIndex)
				dataModelGroup.Post("/single/deleteindex", queryHandlers.DeleteSingleDataModelIndex)
			}
		}
		roleGroup := api.Group("role")
		{
			roleHandlers := new(handlers.RoleHandlers)
			roleGroup.Use(middlewares.FindUserMiddleware())
			roleGroup.Use(middlewares.AuthUserMiddleware())
			roleGroup.Get("/all", roleHandlers.GetAllRoles)
			roleGroup.Post("/add", roleHandlers.AddRole)
			roleGroup.Delete("/:id", roleHandlers.DeleteRole)
			roleGroup.Post("/:id/permission", roleHandlers.UpdateRolePermission)
		}
		settingGroup := api.Group("setting")
		{
			settingHandlers := new(handlers.SettingHandlers)
			settingGroup.Use(middlewares.FindUserMiddleware())
			settingGroup.Use(middlewares.AuthUserMiddleware())
			settingGroup.Get("/single", settingHandlers.GetSingleSetting)
			settingGroup.Post("/single", settingHandlers.UpdateSingleSetting)
		}
		tabGroup := api.Group("tab")
		{
			tabHandlers := new(handlers.TabsHandlers)
			tabGroup.Use(middlewares.FindUserMiddleware())
			tabGroup.Use(middlewares.AuthUserMiddleware())
			tabGroup.Post("/create", tabHandlers.CreateNewTab)
			tabGroup.Post("/update", tabHandlers.UpdateTab)
			tabGroup.Get("/getall/:dbConnId", tabHandlers.GetTabsByDBConnection)
			tabGroup.Delete("/close/:dbConnId/:tabId", tabHandlers.CloseTab)
		}
		consoleGroup := api.Group("console")
		{
			consoleHandlers := new(handlers.ConsoleHandlers)
			consoleGroup.Use(middlewares.FindUserMiddleware())
			consoleGroup.Use(middlewares.AuthUserMiddleware())
			consoleGroup.Post("/runcmd", consoleHandlers.RunCommand)
		}
		aiGroup := api.Group("ai")
		{
			aiHandlers := new(handlers.AIHandlers)
			aiGroup.Use(middlewares.FindUserMiddleware())
			aiGroup.Use(middlewares.AuthUserMiddleware())
			aiGroup.Post("/gensql", aiHandlers.GenerateSQL)
			aiGroup.Get("/listmodels", aiHandlers.ListSupportedAIModels)
		}
	}

	// Serving the Frontend files in Production
	if config.IsLive() {
		app.Use("/assets", filesystem.New(filesystem.Config{
			Root:       http.FS(assets),
			PathPrefix: "frontend/dist/assets",
		}))
		app.Get("/*", func(c *fiber.Ctx) error {
			tokenString := c.Cookies("session")
			if tokenString != "" || c.Path() == "/" {
				file, _ := assets.ReadFile("frontend/dist/index.html")
				c.Set("Content-Type", "text/html; charset=utf-8")
				return c.Send(file)
			}
			return c.Redirect("/", http.StatusTemporaryRedirect)
		})
	}
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"success": true,
		"version": config.GetConfig().Version,
	})
}
