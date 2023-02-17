package controllers

import (
	"encoding/json"
	"errors"

	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/internal/utils"
)

type TabsController struct{}

func (TabsController) CreateTab(dbConnID string) (*models.Tab, error) {

	tab := models.NewBlankTab(dbConnID)

	err := dao.Tab.CreateTab(tab)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return tab, nil
}

func (tc TabsController) GetTabsByDBConnection(dbConnID string) (*[]models.Tab, error) {

	tabs, err := dao.Tab.GetTabsByDBConnectionID(dbConnID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if len(*tabs) == 0 {
		tab, err := tc.CreateTab(dbConnID)
		if err != nil {
			return nil, err
		}
		tabs := []models.Tab{*tab}
		return &tabs, nil
	}

	return tabs, nil
}

func (TabsController) UpdateTab(dbConnID, tabID, tabType string, metadata map[string]interface{}) (*models.Tab, error) {

	if !utils.ContainsString([]string{models.TAB_TYPE_BLANK, models.TAB_TYPE_DATA, models.TAB_TYPE_MODEL, models.TAB_TYPE_HISTORY, models.TAB_TYPE_QUERY}, tabType) {
		return nil, errors.New("invalid tab type")
	}

	metadataStr, err := json.Marshal(metadata)
	if err != nil {
		return nil, errors.New("invalid metadata")
	}

	tab, err := dao.Tab.GetTabByID(dbConnID, tabID)
	if err != nil {
		return nil, errors.New("tab not found")
	}
	tab.Type = tabType
	tab.MetaData = string(metadataStr)

	err = dao.Tab.UpdateTab(dbConnID, tabID, tabType, tab.MetaData)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return tab, nil
}

func (TabsController) CloseTab(dbConnID, tabID string) error {

	err := dao.Tab.DeleteTab(dbConnID, tabID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}
