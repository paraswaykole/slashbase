package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/config"
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
	// api := router.Group("/api/v1")
	// {

	// }
	return router

}
