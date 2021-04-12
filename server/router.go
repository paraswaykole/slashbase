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
			userGroup.POST("/register", userController.RegisterUser)
			userGroup.POST("/verify", userController.VerifySession)
			userGroup.Use(middlewares.FindUserMiddleware())
			userGroup.Use(middlewares.AuthUserMiddleware())
			userGroup.GET("/logout", userController.Logout)
		}
		teamGroup := api.Group("team")
		{
			teamController := new(controllers.TeamController)
			teamGroup.Use(middlewares.FindUserMiddleware())
			teamGroup.Use(middlewares.AuthUserMiddleware())
			teamGroup.POST("/create", teamController.CreateTeam)
			teamGroup.GET("/getall", teamController.GetTeams)
		}
		dbConnGroup := api.Group("dbconnection")
		{
			dbConnController := new(controllers.DBConnectionController)
			dbConnGroup.Use(middlewares.FindUserMiddleware())
			dbConnGroup.Use(middlewares.AuthUserMiddleware())
			dbConnGroup.POST("/create", dbConnController.CreateDBConnection)
			dbConnGroup.GET("/team/:teamId", dbConnController.GetDBConnectionsByTeam)
		}

	}
	return router

}
