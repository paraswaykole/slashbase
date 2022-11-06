package daos

import (
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

type DBConnectionDao struct{}

func (d DBConnectionDao) CreateDBConnection(dbConn *models.DBConnection) error {
	result := db.GetDB().Create(dbConn)
	return result.Error
}

func (d DBConnectionDao) GetDBConnectionsByProject(projectId string) ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ProjectID: projectId}).Find(&dbConns).Error
	return dbConns, err
}

func (d DBConnectionDao) GetDBConnectionsByProjectIds(projectIds []string) ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	sqlQuery := "SELECT * FROM ( SELECT ROW_NUMBER() OVER (PARTITION BY project_id ORDER BY name) AS r, t.* FROM db_connections t where project_id in ?) x WHERE x.r <= 5;"
	err := db.GetDB().Raw(sqlQuery, projectIds).Find(&dbConns).Error
	return dbConns, err
}

func (d DBConnectionDao) GetDBConnectionByID(id string) (*models.DBConnection, error) {
	var dbConn *models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Preload("Project").First(&dbConn).Error
	return dbConn, err
}

func (d DBConnectionDao) DeleteDBConnectionById(id string) error {
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Delete(&models.DBConnection{}).Error
	return err
}
