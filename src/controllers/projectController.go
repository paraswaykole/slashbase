package controllers

import (
	"errors"

	"slashbase.com/backend/src/daos"
	"slashbase.com/backend/src/models"
)

type ProjectController struct{}

var projectDao daos.ProjectDao

func (pc ProjectController) CreateProject(authUser *models.User, projectName string) (*models.Project, *models.ProjectMember, error) {

	if !authUser.IsRoot {
		return nil, nil, errors.New("not allowed")
	}

	project := models.NewProject(authUser, projectName)
	projectMember, err := projectDao.CreateProject(project)
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	return project, projectMember, nil
}

func (pc ProjectController) GetProjects(authUser *models.User) (*[]models.ProjectMember, error) {

	projectMembers, err := projectDao.GetProjectMembersForUser(authUser.ID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (pc ProjectController) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {

	projectMembers, err := projectDao.GetProjectMembers(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (pc ProjectController) AddProjectMembers(projectID, email, role string) (*models.ProjectMember, error) {

	toAddUser, err := userDao.GetUserByEmail(email)
	if err != nil {
		// TODO: Create user and send email if doesn't exist in users table.
		return nil, errors.New("there was some problem")
	}

	newProjectMember, err := models.NewProjectMember(toAddUser.ID, projectID, role)
	if err != nil {
		// TODO: Create user and send email if doesn't exist in users table.
		return nil, err
	}
	err = projectDao.CreateProjectMember(newProjectMember)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	go pc.updateProjectMemberInDBConnections(projectID)
	newProjectMember.User = *toAddUser
	return newProjectMember, nil
}

func (pc ProjectController) updateProjectMemberInDBConnections(projectID string) error {
	dbConns, err := dbConnDao.GetDBConnectionsByProject(projectID)
	if err != nil {
		return err
	}
	dbController := new(DBConnectionController)
	for _, dbConn := range dbConns {
		dbController.updateDBConnUsersProjectMemberRoles(dbConn)
	}
	return nil
}
