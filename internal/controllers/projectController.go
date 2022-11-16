package controllers

import (
	"errors"

	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/models"
)

type ProjectController struct{}

func (pc ProjectController) CreateProject(authUser *models.User, projectName string) (*models.Project, *models.ProjectMember, error) {

	if !authUser.IsRoot {
		return nil, nil, errors.New("not allowed")
	}

	project := models.NewProject(authUser, projectName)
	err := dao.Project.CreateProject(project)
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	role, err := dao.Role.GetAdminRole()
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}

	projectMember := models.NewProjectMember(project.CreatedBy, project.ID, role.ID)
	err = dao.Project.CreateProjectMember(projectMember)
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}
	projectMember.Role = *role

	return project, projectMember, nil
}

func (pc ProjectController) GetProjects(authUser *models.User) (*[]models.ProjectMember, error) {

	projectMembers, err := dao.Project.GetProjectMembersForUser(authUser.ID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (pc ProjectController) DeleteProject(authUser *models.User, id string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, id); err != nil || !isAllowed {
		return err
	}

	project, err := dao.Project.GetProject(id)
	if err != nil {
		return errors.New("could not find project")
	}

	allDBsInProject, err := dao.DBConnection.GetDBConnectionsByProject(project.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	for _, dbConn := range allDBsInProject {
		err := dao.DBConnection.DeleteDBConnectionById(dbConn.ID)
		if err != nil {
			return errors.New("there was some problem deleting db `" + dbConn.Name + "` in the project")
		}
	}

	err = dao.Project.DeleteAllProjectMembersInProject(project.ID)
	if err != nil {
		return errors.New("there was some problem deleting project members")
	}

	err = dao.Project.DeleteProject(project.ID)
	if err != nil {
		return errors.New("there was some problem deleting the project")
	}

	return nil
}

func (pc ProjectController) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {

	projectMembers, err := dao.Project.GetProjectMembers(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (pc ProjectController) AddProjectMember(authUser *models.User, projectID, email, roleID string) (*models.ProjectMember, error) {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectID); err != nil || !isAllowed {
		return nil, err
	}

	toAddUser, err := dao.User.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("there was some problem")
	}

	role, err := dao.Role.GetRoleByID(roleID)
	if err != nil {
		return nil, errors.New("role not found")
	}

	newProjectMember := models.NewProjectMember(toAddUser.ID, projectID, role.ID)
	if err != nil {
		return nil, err
	}
	err = dao.Project.CreateProjectMember(newProjectMember)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	newProjectMember.User = *toAddUser
	newProjectMember.Role = *role
	return newProjectMember, nil
}

func (pc ProjectController) DeleteProjectMember(authUser *models.User, projectId, userId string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectId); err != nil || !isAllowed {
		return err
	}

	projectMember, notFound, err := dao.Project.FindProjectMember(projectId, userId)
	if err != nil {
		if notFound {
			return errors.New("member not found")
		}
		return errors.New("there was some problem")
	}

	err = dao.Project.DeleteProjectMember(projectMember)
	if err != nil {
		return errors.New("there was some problem deleting the member")
	}
	return nil
}

func (pc ProjectController) GetAllRoles(user *models.User) (*[]models.Role, error) {

	if !user.IsRoot {
		return nil, errors.New("not allowed")
	}

	roles, err := dao.Role.GetAllRoles()
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return roles, nil
}

func (pc ProjectController) AddRole(user *models.User, name string) (*models.Role, error) {

	if !user.IsRoot {
		return nil, errors.New("not allowed")
	}

	role := models.NewRole(name)
	err := dao.Role.CreateRole(role)
	if err != nil {
		return nil, errors.New("cannot create role: " + name)
	}
	return role, nil
}

func (pc ProjectController) DeleteRole(user *models.User, roleID string) error {

	if !user.IsRoot {
		return errors.New("not allowed")
	}

	role, err := dao.Role.GetAdminRole()
	if err != nil {
		return errors.New("there was some problem")
	}

	if role.ID == roleID {
		return errors.New("cannot delete admin role")
	}

	err = dao.Role.DeleteRoleByID(roleID)
	if err != nil {
		return errors.New("cannot delete role")
	}

	return nil
}
