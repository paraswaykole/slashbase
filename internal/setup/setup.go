package setup

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

func SetupApp() {
	autoMigrate()
	configureSettings()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(
		&models.Project{},
		&models.DBConnection{},
		&models.DBQuery{},
		&models.DBQueryLog{},
		&models.Setting{},
	)
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
