package server

import (
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/config"
)

// Init server
func Init() {
	nconfig := config.GetConfig()
	if config.IsLive() || config.IsStage() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	router.Run(nconfig.GetString("server.port"))
}
