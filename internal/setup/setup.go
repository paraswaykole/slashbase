package setup

import (
	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/models"
	"gorm.io/gorm"
)

func SetupApp() {
	autoMigrate()
	configureSettings()
}

func autoMigrate() {
	err := db.GetDB().AutoMigrate(
		&models.Project{},
		&models.DBConnection{},
		&models.DBQuery{},
		&models.DBQueryLog{},
		&models.Setting{},
	)
	if err != nil {
		return
	}
}

func configureSettings() {
	_, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_APP_ID)
	if err == gorm.ErrRecordNotFound {
		var settings []models.Setting
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_APP_ID, uuid.New().String()))
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_TELEMETRY_ENABLED, "true"))
		settings = append(settings, *models.NewSetting(models.SETTING_NAME_LOGS_EXPIRE, "30"))
		err := dao.Setting.CreateSettings(&settings)
		if err != nil {
			return
		}
	}
}
