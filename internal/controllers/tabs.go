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

func (TabsController) GetTabsByDBConnection(dbConnectionId string) (*[]models.Tab, error) {

	tabs, err := dao.Tab.GetTabsByDBConnectionID(dbConnectionId)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return tabs, nil
}
