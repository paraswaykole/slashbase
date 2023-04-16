package setup

import (
	"os"

	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/db"
	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"gorm.io/gorm"
)

func SetupServer() {
	autoMigrate()
	configureSettings()
}

func autoMigrate() {
	db.GetDB().AutoMigrate(
		&models.User{},
		&models.UserSession{},
		&common.Project{},
		&models.Role{},
		&models.ProjectMember{},
		&models.RolePermission{},
		&common.DBConnection{},
		&common.DBQuery{},
		&common.DBQueryLog{},
		&common.Tab{},
		&common.Setting{},
	)
	err := db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func configureSettings() {
	_, err := dao.Setting.GetSingleSetting(common.SETTING_NAME_APP_ID)
	if err == gorm.ErrRecordNotFound {
		settings := []common.Setting{}
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_APP_ID, uuid.New().String()))
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_TELEMETRY_ENABLED, "true"))
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_LOGS_EXPIRE, "30"))
		dao.Setting.CreateSettings(&settings)
	}
}
