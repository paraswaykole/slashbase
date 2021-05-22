package views

import "slashbase.com/backend/models"

type DBDataModel struct {
	Name       string `json:"name"`
	SchemaName string `json:"schemaName"`
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
