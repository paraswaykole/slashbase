package dao

import (
	"time"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

type dbQueryLogDao struct{}

var DBQueryLog dbQueryLogDao

func (dbQueryLogDao) CreateDBQueryLog(queryLog *models.DBQueryLog) error {
	err := db.GetDB().Create(queryLog).Error
	return err
}

func (dbQueryLogDao) GetDBQueryLogsDBConnID(dbConnID string, before time.Time) ([]*models.DBQueryLog, error) {
	var dbQueryLogs []*models.DBQueryLog
	err := db.GetDB().Where(&models.DBQueryLog{DBConnectionID: dbConnID}).Where("created_at < ?", before).Preload("User").Order("created_at desc").Limit(config.PAGINATION_COUNT).Find(&dbQueryLogs).Error
	return dbQueryLogs, err
}

func (dbQueryLogDao) ClearOldLogs(days int) error {
	before := time.Now().AddDate(0, 0, -days)
	err := db.GetDB().Where("created_at < ?", before).Delete(&models.DBQueryLog{}).Error
	return err
}
