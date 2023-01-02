package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slashbaseide/slashbase/internal/config"
)

// Init server
func Init() {
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	go router.Run(":" + config.GetServerPort())
}
