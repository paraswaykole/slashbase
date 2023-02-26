package controllers

import (
	"github.com/slashbaseide/slashbase/internal/console"
	"github.com/slashbaseide/slashbase/internal/dao"
)

type ConsoleController struct{}

func (ConsoleController) RunCommand(dbConnectionID, cmdString string) string {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionID)
	if err != nil {
		return "there was some problem"
	}

	return console.HandleCommand(dbConn, cmdString, getQueryConfigsForProjectMember(dbConn))
}
