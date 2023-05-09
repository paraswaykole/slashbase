package dao

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

type dbQueryLogDao struct{}

var DBQueryLog dbQueryLogDao

func (dbQueryLogDao) CreateDBQueryLog(queryLog *models.DBQueryLog) error {
	err := db.GetDB().Create(queryLog).Error
	return err
}

func (dbQueryLogDao) GetDBQueryLogsDBConnID(dbConnID string, before time.Time) ([]*models.DBQueryLog, error) {
	var dbQueryLogs []*models.DBQueryLog
	err := db.GetDB().Model(&models.DBQueryLog{}).Where("db_connection_id = ? AND created_at < ?", dbConnID, before).Order("created_at desc").Limit(config.PAGINATION_COUNT).Preload("User").Find(&dbQueryLogs).Error
	return dbQueryLogs, err
}

func (dbQueryLogDao) ClearOldLogs(days int) error {
	before := time.Now().AddDate(0, 0, -days)
	err := db.GetDB().Where("created_at < ?", before).Delete(&models.DBQueryLog{}).Error
	return err
}
