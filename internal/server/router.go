package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/handlers"
)

// NewRouter return a gin router for server
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "X-Security-Key")
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	api := router.Group("/api/v1")
	api.Use(securityCheck())
	{
		api.GET("health", healthCheck)
		projectGroup := api.Group("project")
		{
			projectHandlers := new(handlers.ProjectHandlers)
			projectGroup.POST("/create", projectHandlers.CreateProject)
			projectGroup.GET("/all", projectHandlers.GetProjects)
			projectGroup.DELETE("/:projectId", projectHandlers.DeleteProject)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnectionHandler := new(handlers.DBConnectionHandlers)
			dbConnGroup.POST("/create", dbConnectionHandler.CreateDBConnection)
			dbConnGroup.GET("/all", dbConnectionHandler.GetDBConnections)
			dbConnGroup.GET("/project/:projectId", dbConnectionHandler.GetDBConnectionsByProject)
			dbConnGroup.GET("/:dbConnId", dbConnectionHandler.GetSingleDBConnection)
			dbConnGroup.DELETE("/:dbConnId", dbConnectionHandler.DeleteDBConnection)
		}
		queryGroup := api.Group("query")
		{
			queryHandlers := new(handlers.QueryHandlers)
			queryGroup.POST("/run", queryHandlers.RunQuery)
			queryGroup.POST("/save/:dbConnId", queryHandlers.SaveDBQuery)
			queryGroup.GET("/getall/:dbConnId", queryHandlers.GetDBQueriesInDBConnection)
			queryGroup.GET("/get/:queryId", queryHandlers.GetSingleDBQuery)
			queryGroup.DELETE("/delete/:queryId", queryHandlers.DeleteDBQuery)
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
				dataModelGroup.POST("/single/addindex", queryHandlers.AddSingleDataModelIndex)
				dataModelGroup.POST("/single/deleteindex", queryHandlers.DeleteSingleDataModelIndex)
			}
		}
		settingGroup := api.Group("setting")
		{
			settingHandlers := new(handlers.SettingHandlers)
			settingGroup.GET("/single", settingHandlers.GetSingleSetting)
			settingGroup.POST("/single", settingHandlers.UpdateSingleSetting)
		}
	}
	return router

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"version": config.GetConfig().Version,
	})
}

func securityCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Security-Key") != config.GetConfig().SecurityKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
