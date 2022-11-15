package controllers

import (
	"errors"

	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/models"
)

type ProjectController struct{}

var projectDao daos.ProjectDao
var roleDao daos.RoleDao

func (pc ProjectController) CreateProject(authUser *models.User, projectName string) (*models.Project, *models.ProjectMember, error) {

	if !authUser.IsRoot {
		return nil, nil, errors.New("not allowed")
	}

	project := models.NewProject(authUser, projectName)
	err := projectDao.CreateProject(project)
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	role, err := roleDao.GetAdminRole()
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	projectMember := models.NewProjectMember(project.CreatedBy, project.ID, role.ID)
	err = projectDao.CreateProjectMember(projectMember)
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

func (pc ProjectController) DeleteProject(authUser *models.User, id string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, id); err != nil || !isAllowed {
		return err
	}

	project, err := projectDao.GetProject(id)
	if err != nil {
		return errors.New("could not find project")
	}

	allDBsInProject, err := dbConnDao.GetDBConnectionsByProject(project.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	for _, dbConn := range allDBsInProject {
		err := dbConnDao.DeleteDBConnectionById(dbConn.ID)
		if err != nil {
			return errors.New("there was some problem deleting db `" + dbConn.Name + "` in the project")
		}
	}

	err = projectDao.DeleteAllProjectMembersInProject(project.ID)
	if err != nil {
		return errors.New("there was some problem deleting project members")
	}

	err = projectDao.DeleteProject(project.ID)
	if err != nil {
		return errors.New("there was some problem deleting the project")
	}

	return nil
}

func (pc ProjectController) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {

	projectMembers, err := projectDao.GetProjectMembers(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (pc ProjectController) AddProjectMember(authUser *models.User, projectID, email, roleID string) (*models.ProjectMember, error) {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectID); err != nil || !isAllowed {
		return nil, err
	}

	toAddUser, err := userDao.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	newProjectMember := models.NewProjectMember(toAddUser.ID, projectID, roleID)
	if err != nil {
		return nil, err
	}
	err = projectDao.CreateProjectMember(newProjectMember)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	newProjectMember.User = *toAddUser
	return newProjectMember, nil
}

func (pc ProjectController) DeleteProjectMember(authUser *models.User, projectId, userId string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectId); err != nil || !isAllowed {
		return err
	}

	projectMember, notFound, err := projectDao.FindProjectMember(projectId, userId)
	if err != nil {
		if notFound {
			return errors.New("member not found")
		}
		return errors.New("there was some problem")
	}

	err = projectDao.DeleteProjectMember(projectMember)
	if err != nil {
		return errors.New("there was some problem deleting the member")
	}
	return nil
}
