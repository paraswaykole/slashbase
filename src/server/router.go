package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/config"
	"slashbase.com/backend/src/middlewares"
	"slashbase.com/backend/src/routes"
)

// NewRouter return a gin router for server
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowCredentials = true
	if config.IsLive() || config.IsDevelopment() {
		corsConfig.AllowOrigins = []string{config.GetAppHost()}
	} else {
		corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	}
	router.Use(cors.New(corsConfig))
	api := router.Group("/api/v1")
	{
		userGroup := api.Group("user")
		{
			userRoutes := new(routes.UserRoutes)
			userGroup.POST("/login", userRoutes.LoginUser)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.POST("/edit", userRoutes.EditAccount)
			userGroup.POST("/add", userRoutes.AddUser)
			userGroup.GET("/all", userRoutes.GetUsers)
			userGroup.GET("/logout", userRoutes.Logout)
		}
		projectGroup := api.Group("project")
		{
			projectRoutes := new(routes.ProjectRoutes)
			projectGroup.Use(middlewares.FindUserMiddleware())
			projectGroup.Use(middlewares.AuthUserMiddleware())
			projectGroup.POST("/create", projectRoutes.CreateProject)
			projectGroup.GET("/all", projectRoutes.GetProjects)
			projectGroup.POST("/:projectId/members/create", projectRoutes.AddProjectMembers)
			projectGroup.GET("/:projectId/members", projectRoutes.GetProjectMembers)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnRoutes := new(routes.DBConnectionRoutes)
			dbConnGroup.Use(middlewares.FindUserMiddleware())
			dbConnGroup.Use(middlewares.AuthUserMiddleware())
			dbConnGroup.POST("/create", dbConnRoutes.CreateDBConnection)
			dbConnGroup.GET("/all", dbConnRoutes.GetDBConnections)
			dbConnGroup.GET("/project/:projectId", dbConnRoutes.GetDBConnectionsByProject)
			dbConnGroup.GET("/:dbConnId", dbConnRoutes.GetSingleDBConnection)
			dbConnGroup.DELETE("/:dbConnId", dbConnRoutes.DeleteDBConnection)
		}
		queryGroup := api.Group("query")
		{
			queryRoutes := new(routes.QueryRoutes)
			queryGroup.Use(middlewares.FindUserMiddleware())
			queryGroup.Use(middlewares.AuthUserMiddleware())
			queryGroup.POST("/run", queryRoutes.RunQuery)
			queryGroup.POST("/save/:dbConnId", queryRoutes.SaveDBQuery)
			queryGroup.GET("/getall/:dbConnId", queryRoutes.GetDBQueriesInDBConnection)
			queryGroup.GET("/get/:queryId", queryRoutes.GetSingleDBQuery)
			queryGroup.GET("/history/:dbConnId", queryRoutes.GetQueryHistoryInDBConnection)
			dataGroup := queryGroup.Group("data")
			{
				dataGroup.GET("/:dbConnId", queryRoutes.GetData)
				dataGroup.POST("/:dbConnId/single", queryRoutes.UpdateSingleData)
				dataGroup.POST("/:dbConnId/add", queryRoutes.AddData)
				dataGroup.POST("/:dbConnId/delete", queryRoutes.DeleteData)
			}
			dataModelGroup := queryGroup.Group("datamodel")
			{
				dataModelGroup.GET("/all/:dbConnId", queryRoutes.GetDataModels)
				dataModelGroup.GET("/single/:dbConnId", queryRoutes.GetSingleDataModel)
			}
		}
	}
	return router

}
