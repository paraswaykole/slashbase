package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/src/config"
)

// Init server
func Init() {
	fmt.Println("Starting slashbase server...")
	if config.IsLive() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := NewRouter()
	router.Run(config.GetServerPort())
}
