package server

import (
	"time"

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
		go func() {
			time.Sleep(500 * time.Millisecond)
			osx.OpenDefault("http://localhost:" + config.GetConfig().Port)
		}()
	}
	router := NewRouter()
	go router.Run(":" + config.GetConfig().Port)
}
