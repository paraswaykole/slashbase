package setup

import (
	"os"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func SetupApp() {
	autoMigrate()
	firstTimeSetup()
	configureRootUser()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(
		&models.User{},
		&models.UserSession{},
		&models.Project{},
		&models.ProjectMember{},
		&models.DBConnection{},
		&models.DBQuery{},
		&models.DBQueryLog{},
	)
	err := db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func firstTimeSetup() {
	// init default settings
}

func configureRootUser() {
	rootUserEmail, rootUserPassword := config.GetRootUser()
	rootUser, err := models.NewUser(rootUserEmail, rootUserPassword)
	if err != nil {
		os.Exit(1)
	}
	rootUser.IsRoot = true
	var userDao daos.UserDao
	_, err = userDao.GetRootUserOrCreate(*rootUser)
	if err != nil {
		os.Exit(1)
	}
}
