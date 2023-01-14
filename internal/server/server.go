package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/icza/gox/osx"
	"github.com/slashbaseide/slashbase/internal/config"
)

// Init server
func Init(isCli bool) {
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.IsLive() {
		go func() {
			time.Sleep(1500 * time.Millisecond)
			osx.OpenDefault("http://localhost:" + config.GetConfig().Port)
		}()
	}
	router := NewRouter()

	if isCli {
		go func() {
			err := router.Run(":" + config.GetConfig().Port)
			if err != nil {
				return
			}
		}()
	} else {
		err := router.Run(":" + config.GetConfig().Port)
		if err != nil {
			return
		}
	}

}
