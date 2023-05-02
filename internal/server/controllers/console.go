package controllers

import (
	"github.com/slashbaseide/slashbase/internal/common/console"
	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type ConsoleController struct{}

func (ConsoleController) RunCommand(authUser *models.User, dbConnectionID, cmdString string) string {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionID)
	if err != nil {
		return "there was some problem"
	}

	pm, err := getIfAuthUserProjectMemberForProject(authUser, dbConn.ProjectID)
	if err != nil {
		return "there was some problem"
	}

	return console.HandleCommand(dbConn, cmdString, getQueryConfigsForProjectMember(pm, dbConn))
}
