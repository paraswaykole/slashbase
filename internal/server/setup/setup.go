package setup

import (
	"os"

	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/common/config"
	commondao "github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/db"
	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"github.com/slashbaseide/slashbase/pkg/ai"
	"gorm.io/gorm"
)

func SetupServer() {
	autoMigrate()
	configureSettings()
	configureRootUser()
	configureRoles()
	initAIClient()
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
		&models.DBQueryLog{},
		&models.Tab{},
		&common.Setting{},
	)
	err := db.GetDB().SetupJoinTable(&models.User{}, "Projects", &models.ProjectMember{})
	if err != nil {
		os.Exit(1)
	}
}

func configureSettings() {
	_, err := commondao.Setting.GetSingleSetting(common.SETTING_NAME_APP_ID)
	if err == gorm.ErrRecordNotFound {
		settings := []common.Setting{}
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_APP_ID, uuid.New().String()))
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_TELEMETRY_ENABLED, "true"))
		settings = append(settings, *common.NewSetting(common.SETTING_NAME_LOGS_EXPIRE, "30"))
		commondao.Setting.CreateSettings(&settings)
	}
}

func configureRootUser() {
	rootUserEmail, rootUserPassword := config.GetConfig().RootUser.Email, config.GetConfig().RootUser.Password
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

func initAIClient() {
	setting, err := commondao.Setting.GetSingleSetting(common.SETTING_NAME_OPENAI_KEY)
	if err == nil {
		ai.InitClient(setting.Value)
	}
}
