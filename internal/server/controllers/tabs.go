package controllers

import (
	"encoding/json"
	"errors"

	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type TabsController struct{}

func (TabsController) CreateTab(authUserID, dbConnID, tabType, modelschema, modelname, queryID, query string) (*models.Tab, error) {

	var tab *common.Tab
	if tabType == common.TAB_TYPE_BLANK {
		tab = common.NewBlankTab(dbConnID)
	} else if tabType == common.TAB_TYPE_DATA {
		tab = common.NewDataTab(dbConnID, modelschema, modelname)
	} else if tabType == common.TAB_TYPE_MODEL {
		tab = common.NewModelTab(dbConnID, modelschema, modelname)
	} else if tabType == common.TAB_TYPE_QUERY {
		tab = common.NewQueryTab(dbConnID, queryID, query)
	} else if tabType == common.TAB_TYPE_HISTORY {
		tab = common.NewHistoryTab(dbConnID)
	} else if tabType == common.TAB_TYPE_CONSOLE {
		tab = common.NewConsoleTab(dbConnID)
	} else if tabType == common.TAB_TYPE_GENSQL {
		tab = common.NewGenSQLTab(dbConnID)
	}

	userTab := models.NewUserTab(tab, authUserID)

	err := dao.Tab.CreateTab(userTab)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return userTab, nil
}

func (tc TabsController) GetTabsByDBConnection(authUserID, dbConnID string) (*[]models.Tab, error) {

	tabs, err := dao.Tab.GetTabsByDBConnectionID(authUserID, dbConnID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if len(*tabs) == 0 {
		tab, err := tc.CreateTab(authUserID, dbConnID, common.TAB_TYPE_BLANK, "", "", "", "")
		if err != nil {
			return nil, err
		}
		tabs := []models.Tab{*tab}
		return &tabs, nil
	}

	return tabs, nil
}

func (TabsController) UpdateTab(authUserID, dbConnID, tabID, tabType string, metadata map[string]interface{}) (*models.Tab, error) {

	if !utils.ContainsString([]string{common.TAB_TYPE_BLANK, common.TAB_TYPE_DATA, common.TAB_TYPE_MODEL, common.TAB_TYPE_HISTORY, common.TAB_TYPE_CONSOLE, common.TAB_TYPE_GENSQL, common.TAB_TYPE_QUERY}, tabType) {
		return nil, errors.New("invalid tab type")
	}

	metadataStr, err := json.Marshal(metadata)
	if err != nil {
		return nil, errors.New("invalid metadata")
	}

	tab, err := dao.Tab.GetTabByID(authUserID, dbConnID, tabID)
	if err != nil {
		return nil, errors.New("tab not found")
	}
	tab.Type = tabType
	tab.MetaData = string(metadataStr)

	err = dao.Tab.UpdateTab(authUserID, dbConnID, tabID, tabType, tab.MetaData)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return tab, nil
}

func (TabsController) CloseTab(authUserID, dbConnID, tabID string) error {

	err := dao.Tab.DeleteTab(authUserID, dbConnID, tabID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}
