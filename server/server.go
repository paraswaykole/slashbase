package server

import (
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/config"
)

// Init server
func Init() {
	if config.IsLive() || config.IsStage() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	router.Run(config.GetServerPort())
}
