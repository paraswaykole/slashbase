package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

type DBConnectionController struct{}

func (DBConnectionController) CreateDBConnection(
	projectID string,
	name string,
	dbtype string,
	scheme string,
	host string,
	port string,
	user string,
	password string,
	dbName string,
	useSSH string,
	sshHost string,
	sshUser string,
	sshPassword string,
	sshKeyFile string,
	useSSL bool,
	isTest bool) (*models.DBConnection, error) {

	dbConn, err := models.NewDBConnection(projectID, name, dbtype, scheme, host, port,
		user, password, dbName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile, useSSL)
	if err != nil {
		return nil, err
	}

	err = queryengines.TestConnection(dbConn.ToQEConnection(), qemodels.NewQueryConfig(false, nil))
	if err != nil {
		return nil, err
	}

	if isTest {
		return nil, nil
	}

	err = dao.DBConnection.CreateDBConnection(dbConn)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConn, nil
}

func (DBConnectionController) GetDBConnections() ([]*models.DBConnection, error) {

	dbConns, err := dao.DBConnection.GetAllDBConnections()
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, err
}

func (DBConnectionController) GetSingleDBConnection(dbConnID string) (*models.DBConnection, error) {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	return dbConn, nil
}

func (DBConnectionController) GetDBConnectionsByProject(projectID string) ([]*models.DBConnection, error) {

	dbConns, err := dao.DBConnection.GetDBConnectionsByProject(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, nil
}

func (DBConnectionController) DeleteDBConnection(dbConnId string) error {
	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return errors.New("db connection not found")
	}

	err = dao.DBConnection.DeleteDBConnectionById(dbConn.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}

func (DBConnectionController) CheckDBConnection(dbConnectionID string) error {

	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnectionID)
	if err != nil {
		return errors.New("there was some problem")
	}

	err = queryengines.TestConnection(dbConn.ToQEConnection(), qemodels.NewQueryConfig(false, nil))
	if err != nil {
		return err
	}

	return err
}
