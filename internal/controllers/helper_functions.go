package controllers

import (
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines/queryconfig"
)

func getQueryConfigsForProjectMember(dbConn *models.DBConnection) *queryconfig.QueryConfig {
	createLog := func(query string) {
		queryLog := models.NewQueryLog(dbConn.ID, query)
		go dao.DBQueryLog.CreateDBQueryLog(queryLog)
	}
	readOnly := false
	return queryconfig.NewQueryConfig(readOnly, createLog)
}
