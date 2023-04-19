package dao

import (
	"github.com/slashbaseide/slashbase/internal/common/db"
	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type tabsDao struct{}

var Tab tabsDao

func (tabsDao) CreateTab(tab *models.Tab) error {
	err := db.GetDB().Create(tab).Error
	return err
}

func (tabsDao) GetTabsByDBConnectionID(userID, dbConnID string) (*[]models.Tab, error) {
	var tabs []models.Tab
	err := db.GetDB().Where(&models.Tab{Tab: common.Tab{DBConnectionID: dbConnID}, UserID: userID}).Find(&tabs).Error
	return &tabs, err
}

func (tabsDao) GetTabByID(userID, dbConnID, tabID string) (*models.Tab, error) {
	var tab models.Tab
	err := db.GetDB().Where(&models.Tab{Tab: common.Tab{ID: tabID, DBConnectionID: dbConnID}, UserID: userID}).First(&tab).Error
	return &tab, err
}

func (tabsDao) UpdateTab(userID, dbConnID, tabID, tabType, metadata string) error {
	err := db.GetDB().Model(&models.Tab{}).Where(&models.Tab{Tab: common.Tab{ID: tabID, DBConnectionID: dbConnID}, UserID: userID}).
		UpdateColumns(map[string]interface{}{"type": tabType, "meta_data": metadata}).Error
	return err
}

func (tabsDao) DeleteTab(userID, dbConnID, tabID string) error {
	err := db.GetDB().Where(models.Tab{Tab: common.Tab{ID: tabID, DBConnectionID: dbConnID}, UserID: userID}).Delete(models.Tab{}).Error
	return err
}
