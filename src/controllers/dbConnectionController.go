package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/models/sbsql"
	"slashbase.com/backend/src/queryengines"
	"slashbase.com/backend/src/utils"
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
	dbUsers, err := dbcc.createRoleLogins(authUser, dbConn)
	if err != nil {
		return nil, err
	}
	dbConn.DBConnectionUsers = append(dbConn.DBConnectionUsers, dbUsers...)

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
	dbConn, err := dbConnDao.GetConnectableRootDBConnection(dbConnId)
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

	go dbcc.deleteRoleLogins(authUser, dbConn)

	return nil
}

func (dbcc DBConnectionController) createRoleLogins(authUser *models.User, dbConn *models.DBConnection) ([]models.DBConnectionUser, error) {
	if dbConn.LoginType != models.DBLOGINTYPE_ROLE_ACCOUNTS {
		return nil, nil
	}

	hasRolePermissions := queryengines.CheckCreateRolePermissions(authUser, dbConn)
	if !hasRolePermissions {
		return nil, fmt.Errorf("user '%s' does not have create role privilege", string(dbConn.ConnectionUser.DBUser))
	}

	dbConnectionUsers := []models.DBConnectionUser{
		{
			DBUser:     sbsql.CryptedData("sb_"+strings.ToLower(models.ROLE_ADMIN)) + "_" + sbsql.CryptedData(strings.ToLower(utils.RandString(4))),
			DBPassword: sbsql.CryptedData(utils.RandString(10)),
			ForRole: sql.NullString{
				String: models.ROLE_ADMIN,
				Valid:  true,
			},
		},
		{
			DBUser:     sbsql.CryptedData("sb_"+strings.ToLower(models.ROLE_DEVELOPER)) + "_" + sbsql.CryptedData(strings.ToLower(utils.RandString(4))),
			DBPassword: sbsql.CryptedData(utils.RandString(10)),
			ForRole: sql.NullString{
				String: models.ROLE_DEVELOPER,
				Valid:  true,
			},
		},
		{
			DBUser:     sbsql.CryptedData("sb_"+strings.ToLower(models.ROLE_ANALYST)) + "_" + sbsql.CryptedData(strings.ToLower(utils.RandString(4))),
			DBPassword: sbsql.CryptedData(utils.RandString(10)),
			ForRole: sql.NullString{
				String: models.ROLE_ANALYST,
				Valid:  true,
			},
		},
	}

	for _, dbUser := range dbConnectionUsers {
		err := queryengines.CreateRoleLogin(authUser, dbConn, &dbUser)
		if err != nil {
			return nil, errors.New("there was some problem")
		}
	}

	projectMembers, err := projectDao.GetProjectMembers(dbConn.ProjectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	for _, projectMember := range *projectMembers {
		if projectMember.Role == models.ROLE_ADMIN {
			dbConnectionUsers[0].UserIDs = append(dbConnectionUsers[0].UserIDs, projectMember.UserID)
		} else if projectMember.Role == models.ROLE_DEVELOPER {
			dbConnectionUsers[1].UserIDs = append(dbConnectionUsers[1].UserIDs, projectMember.UserID)
		} else if projectMember.Role == models.ROLE_ANALYST {
			dbConnectionUsers[2].UserIDs = append(dbConnectionUsers[2].UserIDs, projectMember.UserID)
		}
	}

	return dbConnectionUsers, nil
}

func (dbcc DBConnectionController) deleteRoleLogins(authUser *models.User, dbConn *models.DBConnection) error {

	if dbConn.LoginType != models.DBLOGINTYPE_ROLE_ACCOUNTS {
		return nil
	}

	hasRolePermissions := queryengines.CheckCreateRolePermissions(authUser, dbConn)
	if !hasRolePermissions {
		return fmt.Errorf("user '%s' does not have create role privilege", string(dbConn.ConnectionUser.DBUser))
	}

	for _, dbUser := range dbConn.DBConnectionUsers {
		if dbUser.ForRole.Valid {
			queryengines.DeleteRoleLogin(authUser, dbConn, &dbUser)
		}
	}

	return nil
}

func (dbcc DBConnectionController) updateDBConnUsersProjectMemberRoles(dbConn *models.DBConnection) error {
	if dbConn.LoginType != models.DBLOGINTYPE_ROLE_ACCOUNTS {
		return nil
	}

	dbConnUsers, err := dbConnDao.GetAllRolesDBConnectionUsers(dbConn.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	roleMap := map[string]int{}
	for i, dbUser := range dbConnUsers {
		roleMap[dbUser.ForRole.String] = i
	}

	projectMembers, err := projectDao.GetProjectMembers(dbConn.ProjectID)
	if err != nil {
		return errors.New("there was some problem")
	}

	for role := range roleMap {
		if _, exists := roleMap[role]; exists {
			dbConnUsers[roleMap[role]].UserIDs = []string{}
		}
	}

	for _, projectMember := range *projectMembers {
		dbConnUsers[roleMap[projectMember.Role]].UserIDs = append(dbConnUsers[roleMap[projectMember.Role]].UserIDs, projectMember.UserID)
	}

	for _, dbUser := range dbConnUsers {
		err = dbUser.Save()
		if err != nil {
			return errors.New("there was some problem updating records")
		}
	}
	return nil
}
