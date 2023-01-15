package server

import (
	"io"
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
	router.Use(cors.Default())
	api := router.Group("/api/v1")
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
	router.NoRoute(serveApp)
	return router

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"version": config.GetConfig().Version,
	})
}

func serveApp(c *gin.Context) {
	appUrl := "http://localhost:3000"
	if config.IsLive() {
		appUrl = "https://local.slashbase.com"
	}
	if c.Request.Method == "GET" {
		if resp, err := http.Get(appUrl + c.Request.URL.Path); err == nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				c.String(http.StatusBadGateway, "bad gateway")
				return
			}
			contentType := resp.Header.Get("Content-Type")
			c.Data(http.StatusOK, contentType, body)
			return
		}
		c.String(http.StatusNotFound, "please check your internet connection to load the web IDE.")
		return
	}
}
