package server

import (
	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/config"
)

// Init server
func Init() {
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	router.Run(config.GetServerPort())
}
