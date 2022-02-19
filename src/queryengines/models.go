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
			Name:       tableData["0"].(string),
			SchemaName: tableData["1"].(string),
		}
		return &view
	}
	return nil
}

func BuildDBDataModelField(dbConn *models.DBConnection, fieldData map[string]interface{}) *DBDataModelField {
	if dbConn.Type == models.DBTYPE_POSTGRES {
		view := DBDataModelField{
			Name:       fieldData["0"].(string),
			Type:       fieldData["1"].(string),
			IsNullable: fieldData["2"].(string) == "YES",
		}
		if fieldData["3"] != nil {
			coldef := fieldData["3"].(string)
			view.Default = &coldef
		}
		if fieldData["4"] != nil {
			contype := rune(fieldData["4"].(int8))
			view.IsPrimary = contype == 'p'
		}
		if fieldData["5"] != nil {
			maxLen := fieldData["5"].(int32)
			view.CharMaxLength = &maxLen
		}

		return &view
	}
	return nil
}
