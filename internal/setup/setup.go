package setup

import (
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func SetupApp() {
	autoMigrate()
	configureSettings()
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
		&models.Setting{},
	)
	err := db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func configureSettings() {
	var settingsDao daos.SettingDao
	_, err := settingsDao.GetSingleSetting(models.SETTING_NAME_APP_ID)
	if err == gorm.ErrRecordNotFound {
		settings := []models.Setting{}
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_APP_ID, uuid.New().String()))
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_TELEMETRY_ENABLED, "true"))
		settingsDao.CreateSettings(&settings)
	}
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
