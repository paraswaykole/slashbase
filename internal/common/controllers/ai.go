package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/pkg/ai"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

type AIController struct{}

func (AIController) GenerateSQL(dbConnectionID, text string) (string, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionID)
	if err != nil {
		return "", errors.New("there was some problem")
	}

	datamodels, err := queryengines.GetDataModels(dbConn.ToQEConnection(), qemodels.NewQueryConfig(false, nil))
	if err != nil {
		return "", errors.New("there was some problem fetching data models")
	}

	for i, dm := range datamodels {
		dm, err := queryengines.GetSingleDataModel(dbConn.ToQEConnection(), dm.SchemaName, dm.Name, qemodels.NewQueryConfig(false, nil))
		if err == nil {
			datamodels[i] = dm
		}
	}

	return ai.GenerateSQL(dbConn.Type, text, datamodels)
}

func (AIController) GetModels() []string {
	return ai.ListSupportedOpenAiModels()
}
