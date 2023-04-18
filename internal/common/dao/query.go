package dao

import (
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/common/models"
)

type dbQueryDao struct{}

var DBQuery dbQueryDao

func (dbQueryDao) CreateQuery(query *models.DBQuery) error {
	err := db.GetDB().Create(query).Error
	return err
}

func (dbQueryDao) GetDBQueriesByDBConnId(dbConnID string) ([]*models.DBQuery, error) {
	var dbQueries []*models.DBQuery
	err := db.GetDB().Where(&models.DBQuery{DBConnectionID: dbConnID}).Find(&dbQueries).Error
	return dbQueries, err
}

func (dbQueryDao) GetSingleDBQuery(queryID string) (*models.DBQuery, error) {
	var dbQuery models.DBQuery
	err := db.GetDB().Where(&models.DBQuery{ID: queryID}).Preload("DBConnection").First(&dbQuery).Error
	return &dbQuery, err
}

func (dbQueryDao) UpdateDBQuery(queryID string, dbQuery *models.DBQuery) error {
	err := db.GetDB().Where(&models.DBQuery{ID: queryID}).Updates(dbQuery).Error
	return err
}

func (dbQueryDao) DeleteDBQuery(queryID string) error {
	err := db.GetDB().Where(models.DBQuery{ID: queryID}).Delete(models.DBQuery{}).Error
	return err
}
