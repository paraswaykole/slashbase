package dao

import (
	"github.com/slashbaseide/slashbase/internal/db"
	"github.com/slashbaseide/slashbase/internal/models"
)

type dbConnectionDao struct{}

var DBConnection dbConnectionDao

func (dbConnectionDao) CreateDBConnection(dbConn *models.DBConnection) error {
	result := db.GetDB().Create(dbConn)
	return result.Error
}

func (dbConnectionDao) GetDBConnectionsByProject(projectId string) ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ProjectID: projectId}).Find(&dbConns).Error
	return dbConns, err
}

func (dbConnectionDao) GetAllDBConnections() ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	err := db.GetDB().Model(&models.DBConnection{}).Find(&dbConns).Error
	return dbConns, err
}

func (dbConnectionDao) GetAllDBConnectionsCount() (int64, error) {
	var count int64
	err := db.GetDB().Model(models.DBConnection{}).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (dbConnectionDao) GetDBConnectionByID(id string) (*models.DBConnection, error) {
	var dbConn *models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Preload("Project").First(&dbConn).Error
	return dbConn, err
}

func (dbConnectionDao) DeleteDBConnectionById(id string) error {
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Delete(&models.DBConnection{}).Error
	return err
}

func (dbConnectionDao) GetDBConnectionByName(name string) (*models.DBConnection, error) {
	var dbConn *models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{Name: name}).Preload("Project").First(&dbConn).Error
	return dbConn, err
}
