package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
)

type TabsController struct{}

func (TabsController) CreateTab(dbConnectionId string) (*models.Tab, error) {

	tab := models.NewBlankTab(dbConnectionId)

	err := dao.Tab.CreateTab(tab)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return tab, nil
}

func (tc TabsController) GetTabsByDBConnection(dbConnectionId string) (*[]models.Tab, error) {

	tabs, err := dao.Tab.GetTabsByDBConnectionID(dbConnectionId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	if len(*tabs) == 0 {
		tab, err := tc.CreateTab(dbConnectionId)
		if err != nil {
			return nil, err
		}
		tabs := []models.Tab{*tab}
		return &tabs, nil
	}

	return tabs, nil
}
