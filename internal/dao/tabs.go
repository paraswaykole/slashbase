package dao

import (
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/models"
)

type tabsDao struct{}

var Tab tabsDao

func (tabsDao) CreateTab(tab *models.Tab) error {
	err := db.GetDB().Create(tab).Error
	return err
}

func (tabsDao) GetTabsByDBConnectionID(dbConnID string) (*[]models.Tab, error) {
	var tabs []models.Tab
	err := db.GetDB().Where(&models.Tab{DBConnectionID: dbConnID}).Find(&tabs).Error
	return &tabs, err
}

func (tabsDao) GetTabByID(dbConnID, tabID string) (*models.Tab, error) {
	var tab models.Tab
	err := db.GetDB().Where(&models.Tab{ID: tabID, DBConnectionID: dbConnID}).First(&tab).Error
	return &tab, err
}

func (tabsDao) UpdateTab(dbConnID, tabID, tabType, metadata string) error {
	err := db.GetDB().Model(&models.Tab{}).Where(&models.Tab{ID: tabID, DBConnectionID: dbConnID}).
		UpdateColumns(map[string]interface{}{"type": tabType, "meta_data": metadata}).Error
	return err
}

func (tabsDao) DeleteTab(dbConnID, tabID string) error {
	err := db.GetDB().Where(models.Tab{ID: tabID, DBConnectionID: dbConnID}).Delete(models.Tab{}).Error
	return err
}
