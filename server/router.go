package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/config"
	"slashbase.com/backend/controllers"
	"slashbase.com/backend/middlewares"
)

// NewRouter return a gin router for server
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowCredentials = true
	if config.IsLive() {
		corsConfig.AllowOrigins = []string{"https://slashbase.com", "https://www.slashbase.com"}
	} else if config.IsStage() {
		corsConfig.AllowOrigins = []string{"https://staging.slashbase.com"}
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	}
	router.Use(cors.New(corsConfig))
	api := router.Group("/api/v1")
	{
		userGroup := api.Group("user")
		{
			userController := new(controllers.UserController)
			userGroup.POST("/login", userController.LoginUser)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.POST("/add", userController.AddUser)
			userGroup.GET("/logout", userController.Logout)
		}
		projectGroup := api.Group("project")
		{
			projectController := new(controllers.ProjectController)
			projectGroup.Use(middlewares.FindUserMiddleware())
			projectGroup.Use(middlewares.AuthUserMiddleware())
			projectGroup.POST("/create", projectController.CreateProject)
			projectGroup.GET("/all", projectController.GetProjects)
			projectGroup.GET("/members/:projectId", projectController.GetProjectMembers)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnController := new(controllers.DBConnectionController)
			dbConnGroup.Use(middlewares.FindUserMiddleware())
			dbConnGroup.Use(middlewares.AuthUserMiddleware())
			dbConnGroup.POST("/create", dbConnController.CreateDBConnection)
			dbConnGroup.GET("/all", dbConnController.GetDBConnections)
			dbConnGroup.GET("/project/:projectId", dbConnController.GetDBConnectionsByProject)
			dbConnGroup.GET("/:dbConnId", dbConnController.GetSingleDBConnection)
		}
		queryGroup := api.Group("query")
		{
			queryController := new(controllers.QueryController)
			queryGroup.Use(middlewares.FindUserMiddleware())
			queryGroup.Use(middlewares.AuthUserMiddleware())
			queryGroup.POST("/run", queryController.RunQuery)
			queryGroup.GET("/datamodels/:dbConnId", queryController.GetDataModels)
		}
	}
	return router

}
