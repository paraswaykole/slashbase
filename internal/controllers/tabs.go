package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
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

func (TabsController) CloseTab(dbConnID, tabID string) error {

	err := dao.Tab.DeleteTab(dbConnID, tabID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}
