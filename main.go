package main

import (
	"os"
	"path/filepath"

	"github.com/slashbaseide/slashbase/cmd"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/server"
	"github.com/slashbaseide/slashbase/internal/setup"
	"github.com/slashbaseide/slashbase/internal/tasks"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
)

var build = config.BUILD_DEVELOPMENT
var version = config.BUILD_DEVELOPMENT

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath, _ := filepath.EvalSymlinks(ex)
	exPath = filepath.Dir(exPath)
	config.Init(exPath, build, version)
	db.InitGormDB(exPath)
	setup.SetupApp()
	queryengines.Init()
	tasks.InitCron()
	server.Init()
	cmd.Execute()
}
