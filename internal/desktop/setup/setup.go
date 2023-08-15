package setup

import (
	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/pkg/ai"
	"gorm.io/gorm"
)

func SetupApp() {
	autoMigrate()
	configureSettings()
	initAIClient()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(
		&models.Project{},
		&models.DBConnection{},
		&models.Tab{},
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

func initAIClient() {
	setting, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_OPENAI_KEY)
	if err == nil {
		ai.InitClient(setting.Value)
		setting, _ := dao.Setting.GetSingleSetting(models.SETTING_NAME_OPENAI_MODEL)
		ai.SetOpenAiModel(setting.Value)
	}
}
