package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

type DBConnectionDao struct{}

func (d DBConnectionDao) CreateDBConnection(dbConn *models.DBConnection) error {
	result := db.GetDB().Create(dbConn)
	return result.Error
}

func (d DBConnectionDao) GetDBConnectionsByTeam(teamId string) ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{TeamID: teamId}).Find(&dbConns).Error
	return dbConns, err
}

func (d DBConnectionDao) GetDBConnectionsByTeamIds(teamIds []string) ([]*models.DBConnection, error) {
	var dbConns []*models.DBConnection
	err := db.GetDB().Where("team_id IN ?", teamIds).Find(&dbConns).Error
	return dbConns, err
}

func (d DBConnectionDao) GetDBConnectionByID(id string) (*models.DBConnection, error) {
	var dbConn *models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Preload("Team").First(&dbConn).Error
	return dbConn, err
}
