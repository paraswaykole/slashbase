package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slashbaseide/slashbase/internal/config"
)

// Init server
func Init() {
	fmt.Println("Connect to Slashbase IDE at https://app.slashbase.com")
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	go router.Run(":" + config.GetServerPort())
}
