package controllers

import (
	"errors"

	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/pkg/queryengines"
)

type DBConnectionController struct{}

var dbConnDao daos.DBConnectionDao

func (dbcc DBConnectionController) CreateDBConnection(
	authUser *models.User,
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
	sshKeyFile string) (*models.DBConnection, error) {

	dbConn, err := models.NewDBConnection(authUser.ID, projectID, name, dbtype, scheme, host, port,
		user, password, dbName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile)
	if err != nil {
		return nil, err
	}

	dbConnCopy := *dbConn
	success := queryengines.TestConnection(authUser, &dbConnCopy)
	if !success {
		return nil, errors.New("failed to connect to database")
	}

	err = dbConnDao.CreateDBConnection(dbConn)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConn, nil
}

func (dbcc DBConnectionController) GetDBConnections(authUserProjectIds *[]string) ([]*models.DBConnection, error) {

	dbConns, err := dbConnDao.GetDBConnectionsByProjectIds(*authUserProjectIds)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, err
}

func (dbcc DBConnectionController) GetSingleDBConnection(authUser *models.User, dbConnID string) (*models.DBConnection, error) {
	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	// TODO: check if authUser is member of project
	return dbConn, nil
}

func (dbcc DBConnectionController) GetDBConnectionsByProject(projectID string) ([]*models.DBConnection, error) {

	dbConns, err := dbConnDao.GetDBConnectionsByProject(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, nil
}

func (dbcc DBConnectionController) DeleteDBConnection(authUser *models.User, dbConnId string) error {
	dbConn, err := dbConnDao.GetDBConnectionByID(dbConnId)
	if err != nil {
		return errors.New("db connection not found")
	}

	if _, err := GetAuthUserHasRolesForProject(authUser, dbConn.ProjectID, []string{models.ROLE_ADMIN}); err != nil {
		return err
	}

	err = dbConnDao.DeleteDBConnectionById(dbConn.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}
