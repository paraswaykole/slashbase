package controllers

import (
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

func getQueryConfigsForProjectMember(dbConn *models.DBConnection) *qemodels.QueryConfig {
	createLog := func(query string) {
		queryLog := models.NewQueryLog(dbConn.ID, query)
		go dao.DBQueryLog.CreateDBQueryLog(queryLog)
	}
	readOnly := false
	return qemodels.NewQueryConfig(readOnly, createLog)
}
