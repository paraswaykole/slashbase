package controllers

import (
	"errors"

	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/pkg/queryengines"
	"slashbase.com/backend/pkg/queryengines/queryconfig"
)

type DBConnectionController struct{}

func (DBConnectionController) CreateDBConnection(
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

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectID); err != nil || !isAllowed {
		return nil, err
	}

	dbConn, err := models.NewDBConnection(authUser.ID, projectID, name, dbtype, scheme, host, port,
		user, password, dbName, useSSH, sshHost, sshUser, sshPassword, sshKeyFile)
	if err != nil {
		return nil, err
	}

	dbConnCopy := *dbConn
	success := queryengines.TestConnection(authUser, &dbConnCopy, queryconfig.NewQueryConfig(true, nil))
	if !success {
		return nil, errors.New("failed to connect to database")
	}

	err = dao.DBConnection.CreateDBConnection(dbConn)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConn, nil
}

func (DBConnectionController) GetDBConnections(authUserProjectIds *[]string) ([]*models.DBConnection, error) {

	dbConns, err := dao.DBConnection.GetDBConnectionsByProjectIds(*authUserProjectIds)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, err
}

func (DBConnectionController) GetSingleDBConnection(authUser *models.User, dbConnID string) (*models.DBConnection, error) {
	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	// TODO: check if authUser is member of project
	return dbConn, nil
}

func (DBConnectionController) GetDBConnectionsByProject(projectID string) ([]*models.DBConnection, error) {

	dbConns, err := dao.DBConnection.GetDBConnectionsByProject(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return dbConns, nil
}

func (DBConnectionController) DeleteDBConnection(authUser *models.User, dbConnId string) error {
	dbConn, err := dao.DBConnection.GetDBConnectionByID(dbConnId)
	if err != nil {
		return errors.New("db connection not found")
	}

	if _, err := getAuthUserHasAdminRoleForProject(authUser, dbConn.ProjectID); err != nil {
		return err
	}

	err = dao.DBConnection.DeleteDBConnectionById(dbConn.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}
