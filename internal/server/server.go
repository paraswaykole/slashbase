package server

import (
	"github.com/gin-gonic/gin"
	"github.com/icza/gox/osx"
	"github.com/slashbaseide/slashbase/internal/config"
)

// Init server
func Init() {
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.IsLive() {
		osx.OpenDefault("https://app.slashbase.com/local")
	}
	router := NewRouter()
	go router.Run(":" + config.GetServerPort())
}
