package dao

import (
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"gorm.io/gorm/clause"
)

type settingDao struct{}

var Setting settingDao

func (settingDao) CreateSetting(setting *models.Setting) error {
	err := db.GetDB().Create(setting).Error
	return err
}

func (settingDao) CreateSettings(settings *[]models.Setting) error {
	err := db.GetDB().Create(settings).Error
	return err
}

func (settingDao) GetSingleSetting(name string) (*models.Setting, error) {
	var setting models.Setting
	err := db.GetDB().Where(&models.Setting{Name: name}).First(&setting).Error
	return &setting, err
}

func (settingDao) UpdateSingleSetting(name, value string) error {
	setting := models.NewSetting(name, value)
	err := db.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(setting).Error
	return err
}
