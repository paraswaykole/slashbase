package views

import "slashbase.com/backend/models"

type DBDataModel struct {
	Name       string `json:"name"`
	SchemaName string `json:"schemaName"`
}

type DBQueryView struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Query          string `json:"query"`
	DBConnectionID string `json:"dbConnectionId"`
}

func BuildDBDataModel(dbConnection *models.DBConnection, tableData map[string]interface{}) *DBDataModel {
	if dbConnection.Type == models.DBTYPE_POSTGRES {
		view := DBDataModel{
			Name:       tableData["tablename"].(string),
			SchemaName: tableData["schemaname"].(string),
		}
		return &view
	}
	return nil
}

func BuildDBQueryView(query *models.DBQuery) *DBQueryView {
	return &DBQueryView{
		ID:             query.ID,
		Name:           query.Name,
		Query:          query.Query,
		DBConnectionID: query.DBConnectionID,
	}
}
