package views

import "slashbase.com/backend/models"

type DBTableView struct {
	TableName  string `json:"tableName"`
	SchemaName string `json:"schemaName"`
}

func BuildDBTableView(dbConnection *models.DBConnection, tableData map[string]interface{}) *DBTableView {
	if dbConnection.Type == models.DBTYPE_POSTGRES {
		view := DBTableView{
			TableName:  tableData["tablename"].(string),
			SchemaName: tableData["schemaname"].(string),
		}
		return &view
	}
	return nil
}
