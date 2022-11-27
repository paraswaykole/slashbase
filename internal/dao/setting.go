package dao

import (
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
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
	err := db.GetDB().Model(models.Setting{}).Where(&models.Setting{Name: name}).Update("value", value).Error
	return err
}
