package controllers

import (
	"errors"

	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/queryengines"
)

type DBConnectionController struct{}

var dbConnDao daos.DBConnectionDao

func (dbcc DBConnectionController) CreateDBConnection(
	authUser *models.User,
	projectID string,
	name string,
	host string,
	port string,
	password string,
	user string,
	dbName string,
	loginType string,
	useSSH string,
	sshHost string,
	sshUser string,
	sshPassword string,
	sshKeyFile string) (*models.DBConnection, error) {

	dbConn, err := models.NewPostgresDBConnection(authUser.ID, projectID, name, host, port,
		user, password, dbName, loginType, useSSH, sshHost, sshUser, sshPassword, sshKeyFile)
	if err != nil {
		return nil, err
	}

	success := queryengines.TestConnection(authUser, dbConn)
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
