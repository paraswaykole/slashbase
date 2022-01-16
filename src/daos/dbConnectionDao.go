package daos

import (
	"slashbase.com/backend/src/db"
	"slashbase.com/backend/src/models"
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

func (d DBConnectionDao) GetConnectableDBConnection(id, userID string) (*models.DBConnection, error) {
	var dbConn *models.DBConnection
	err := db.GetDB().Where(&models.DBConnection{ID: id}).Preload("Project").First(&dbConn).Error
	if err == nil {
		var dbConnUser models.DBConnectionUser
		if dbConn.LoginType == models.DBLOGINTYPE_ROOT {
			err = db.GetDB().Where("db_connection_id = ? AND is_root = ?", id, true).First(&dbConnUser).Error
		} else {
			err = db.GetDB().Where("db_connection_id = ? AND ? = ANY(user_ids)", id, userID).First(&dbConnUser).Error
		}
		if err != nil {
			return nil, err
		}
		dbConn.ConnectionUser = &dbConnUser
	}
	return dbConn, err
}

func (d DBConnectionDao) GetAllRolesDBConnectionUsers(dbConnectionID string) ([]*models.DBConnectionUser, error) {
	var dbConnUsers []*models.DBConnectionUser
	err := db.GetDB().Where("db_connection_id = ? AND for_role IN ?", dbConnectionID, []string{models.ROLE_ADMIN, models.ROLE_DEVELOPER, models.ROLE_ANALYST}).Find(&dbConnUsers).Error
	if err != nil {
		return nil, err
	}
	return dbConnUsers, nil
}
