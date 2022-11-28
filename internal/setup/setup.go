package setup

import (
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func SetupApp() {
	autoMigrate()
	configureSettings()
	configureRootUser()
	configureRoles()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(
		&models.User{},
		&models.UserSession{},
		&models.Project{},
		&models.Role{},
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
	_, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_APP_ID)
	if err == gorm.ErrRecordNotFound {
		settings := []models.Setting{}
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_APP_ID, uuid.New().String()))
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_TELEMETRY_ENABLED, "true"))
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_LOGS_EXPIRE, "30"))
		dao.Setting.CreateSettings(&settings)
	}
}

func configureRootUser() {
	rootUserEmail, rootUserPassword := config.GetRootUser()
	rootUser, err := models.NewUser(rootUserEmail, rootUserPassword)
	if err != nil {
		os.Exit(1)
	}
	rootUser.IsRoot = true
	_, err = dao.User.GetRootUserOrCreate(*rootUser)
	if err != nil {
		os.Exit(1)
	}
}

func configureRoles() {
	_, err := dao.Role.GetAdminRole()
	if err != nil {
		os.Exit(1)
	}
}
