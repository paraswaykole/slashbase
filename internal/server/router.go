package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/handlers"
	"slashbase.com/backend/internal/middlewares"
)

// NewRouter return a gin router for server
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowCredentials = true
	corsConfig.AllowOriginFunc = func(origin string) bool {
		return true
	}
	router.Use(cors.New(corsConfig))
	api := router.Group("/api/v1")
	{
		api.GET("health", healthCheck)
		userGroup := api.Group("user")
		{
			userHandlers := new(handlers.UserHandlers)
			userGroup.POST("/login", userHandlers.LoginUser)
			userGroup.GET("/checkauth", userHandlers.CheckAuth)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.POST("/edit", userHandlers.EditAccount)
			userGroup.POST("/password", userHandlers.ChangePassword)
			userGroup.POST("/add", userHandlers.AddUser)
			userGroup.GET("/all", userHandlers.GetUsers)
			userGroup.GET("/logout", userHandlers.Logout)
		}
		projectGroup := api.Group("project")
		{
			projectHandlers := new(handlers.ProjectHandlers)
			projectGroup.Use(middlewares.FindUserMiddleware())
			projectGroup.Use(middlewares.AuthUserMiddleware())
			projectGroup.POST("/create", projectHandlers.CreateProject)
			projectGroup.GET("/all", projectHandlers.GetProjects)
			projectGroup.DELETE("/:projectId", projectHandlers.DeleteProject)
			projectGroup.POST("/:projectId/members/create", projectHandlers.AddProjectMember)
			projectGroup.DELETE("/:projectId/members/:userId", projectHandlers.DeleteProjectMember)
			projectGroup.GET("/:projectId/members", projectHandlers.GetProjectMembers)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnectionHandler := new(handlers.DBConnectionHandlers)
			dbConnGroup.Use(middlewares.FindUserMiddleware())
			dbConnGroup.Use(middlewares.AuthUserMiddleware())
			dbConnGroup.POST("/create", dbConnectionHandler.CreateDBConnection)
			dbConnGroup.GET("/all", dbConnectionHandler.GetDBConnections)
			dbConnGroup.GET("/project/:projectId", dbConnectionHandler.GetDBConnectionsByProject)
			dbConnGroup.GET("/:dbConnId", dbConnectionHandler.GetSingleDBConnection)
			dbConnGroup.DELETE("/:dbConnId", dbConnectionHandler.DeleteDBConnection)
		}
		queryGroup := api.Group("query")
		{
			queryHandlers := new(handlers.QueryHandlers)
			queryGroup.Use(middlewares.FindUserMiddleware())
			queryGroup.Use(middlewares.AuthUserMiddleware())
			queryGroup.POST("/run", queryHandlers.RunQuery)
			queryGroup.POST("/save/:dbConnId", queryHandlers.SaveDBQuery)
			queryGroup.GET("/getall/:dbConnId", queryHandlers.GetDBQueriesInDBConnection)
			queryGroup.GET("/get/:queryId", queryHandlers.GetSingleDBQuery)
			queryGroup.GET("/history/:dbConnId", queryHandlers.GetQueryHistoryInDBConnection)
			dataGroup := queryGroup.Group("data")
			{
				dataGroup.GET("/:dbConnId", queryHandlers.GetData)
				dataGroup.POST("/:dbConnId/single", queryHandlers.UpdateSingleData)
				dataGroup.POST("/:dbConnId/add", queryHandlers.AddData)
				dataGroup.POST("/:dbConnId/delete", queryHandlers.DeleteData)
			}
			dataModelGroup := queryGroup.Group("datamodel")
			{
				dataModelGroup.GET("/all/:dbConnId", queryHandlers.GetDataModels)
				dataModelGroup.GET("/single/:dbConnId", queryHandlers.GetSingleDataModel)
				dataModelGroup.POST("/single/addfield", queryHandlers.AddSingleDataModelField)
				dataModelGroup.POST("/single/deletefield", queryHandlers.DeleteSingleDataModelField)
			}
		}
		settingGroup := api.Group("setting")
		{
			settingHandlers := new(handlers.SettingHandlers)
			settingGroup.Use(middlewares.FindUserMiddleware())
			settingGroup.Use(middlewares.AuthUserMiddleware())
			settingGroup.GET("/single", settingHandlers.GetSingleSetting)
			settingGroup.POST("/single", settingHandlers.UpdateSingleSetting)
		}
		roleGroup := api.Group("role")
		{
			roleHandlers := new(handlers.RoleHandlers)
			roleGroup.Use(middlewares.FindUserMiddleware())
			roleGroup.Use(middlewares.AuthUserMiddleware())
			roleGroup.GET("/all", roleHandlers.GetAllRoles)
			roleGroup.POST("/add", roleHandlers.AddRole)
			roleGroup.DELETE("/:id", roleHandlers.DeleteRole)
		}
	}

	// Serving the Frontend files in Production
	if config.IsLive() {
		router.LoadHTMLGlob("html/*.html")
		router.Static("_next", "html/_next")
		router.StaticFile("favicon.ico", "html/favicon.ico")
		router.StaticFile("logo-icon.svg", "html/logo-icon.svg")
		router.StaticFile("logo.svg", "html/logo.svg")
		router.NoRoute(func(c *gin.Context) {
			tokenString, _ := c.Cookie("session")
			if tokenString != "" || c.Request.URL.Path == "/login" {
				if c.Request.URL.Path == "/login" && tokenString != "" {
					c.Redirect(http.StatusTemporaryRedirect, "/home")
					return
				}
				c.HTML(http.StatusOK, "index.html", nil)
				return
			}
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		})
	}
	return router

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"version": config.VERSION,
	})
}
