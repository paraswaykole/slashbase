package daos

import (
	"slashbase.com/backend/src/db"
	"slashbase.com/backend/src/models"
)

type DBQueryDao struct{}

func (d DBQueryDao) CreateQuery(query *models.DBQuery) error {
	err := db.GetDB().Create(query).Error
	return err
}

func (d DBQueryDao) GetDBQueriesByDBConnId(dbConnID string) ([]*models.DBQuery, error) {
	var dbQueries []*models.DBQuery
	err := db.GetDB().Where(&models.DBQuery{DBConnectionID: dbConnID}).Find(&dbQueries).Error
	return dbQueries, err
}

func (d DBQueryDao) GetSingleDBQuery(queryID string) (*models.DBQuery, error) {
	var dbQuery models.DBQuery
	err := db.GetDB().Where(&models.DBQuery{ID: queryID}).Preload("DBConnection").First(&dbQuery).Error
	return &dbQuery, err
}

func (d DBQueryDao) UpdateDBQuery(queryID string, dbQuery *models.DBQuery) error {
	err := db.GetDB().Where(&models.DBQuery{ID: queryID}).Updates(dbQuery).Error
	return err
}
