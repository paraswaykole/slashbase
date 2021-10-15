package queryengines

import (
	"slashbase.com/backend/src/models"
)

type DBDataModel struct {
	Name       string             `json:"name"`
	SchemaName string             `json:"schemaName"`
	Fields     []DBDataModelField `json:"fields"`
}

type DBDataModelField struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	IsPrimary     bool    `json:"isPrimary"`
	IsNullable    bool    `json:"isNullable"`
	CharMaxLength *int32  `json:"charMaxLength"`
	Default       *string `json:"default"`
}

func BuildDBDataModel(dbConn *models.DBConnection, tableData map[string]interface{}) *DBDataModel {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		view := DBDataModel{
			Name:       tableData["tablename"].(string),
			SchemaName: tableData["schemaname"].(string),
		}
		return &view
	}
	return nil
}

func BuildDBDataModelField(dbConn *models.DBConnection, fieldData map[string]interface{}) *DBDataModelField {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		view := DBDataModelField{
			Name:       fieldData["column_name"].(string),
			Type:       fieldData["data_type"].(string),
			IsNullable: fieldData["is_nullable"].(string) == "YES",
		}
		if fieldData["column_default"] != nil {
			coldef := fieldData["column_default"].(string)
			view.Default = &coldef
		}
		if fieldData["contype"] != nil {
			contype := rune(fieldData["contype"].(int8))
			view.IsPrimary = contype == 'p'
		}
		if fieldData["character_maximum_length"] != nil {
			maxLen := fieldData["character_maximum_length"].(int32)
			view.CharMaxLength = &maxLen
		}

		return &view
	}
	return nil
}
