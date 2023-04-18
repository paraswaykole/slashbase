package controllers

import (
	"errors"

	commonctrlr "github.com/slashbaseide/slashbase/internal/common/controllers"
	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

var commonProjectController commonctrlr.ProjectController

type ProjectController struct{}

func (ProjectController) CreateProject(authUser *models.User, projectName string) (*common.Project, *models.ProjectMember, error) {

	project, err := commonProjectController.CreateProject(projectName)
	if err != nil {
		return nil, nil, err
	}

	role, err := dao.Role.GetAdminRole()
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}
	rolePermissions, _ := dao.RolePermission.GetRolePermissionsForRole(role.ID)
	role.Permissions = *rolePermissions

	projectMember := models.NewProjectMember(authUser.ID, project.ID, role.ID)
	projectMember.IsCreator = true
	err = dao.ProjectMember.CreateProjectMember(projectMember)
	if err != nil {
		return nil, nil, errors.New("there was some problem")
	}
	projectMember.Role = *role

	return project, projectMember, nil
}

func (ProjectController) GetProjects(authUser *models.User) (*[]models.ProjectMember, error) {

	projectMembers, err := dao.ProjectMember.GetProjectMembersForUser(authUser.ID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (ProjectController) DeleteProject(authUser *models.User, id string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, id); err != nil || !isAllowed {
		return err
	}

	err := dao.ProjectMember.DeleteAllProjectMembersInProject(id)
	if err != nil {
		return errors.New("there was some problem deleting project members")
	}

	err = commonProjectController.DeleteProject(id)
	if err != nil {
		return errors.New("there was some problem deleting the project")
	}

	return nil
}

func (ProjectController) GetProjectMembers(projectID string) (*[]models.ProjectMember, error) {

	projectMembers, err := dao.ProjectMember.GetProjectMembers(projectID)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projectMembers, nil
}

func (ProjectController) AddProjectMember(authUser *models.User, projectID, email, roleID string) (*models.ProjectMember, error) {

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
	err = dao.ProjectMember.CreateProjectMember(newProjectMember)
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	newProjectMember.User = *toAddUser
	newProjectMember.Role = *role
	return newProjectMember, nil
}

func (ProjectController) DeleteProjectMember(authUser *models.User, projectId, userId string) error {

	if isAllowed, err := getAuthUserHasAdminRoleForProject(authUser, projectId); err != nil || !isAllowed {
		return err
	}

	projectMember, notFound, err := dao.ProjectMember.FindProjectMember(projectId, userId)
	if err != nil {
		if notFound {
			return errors.New("member not found")
		}
		return errors.New("there was some problem")
	}

	err = dao.ProjectMember.DeleteProjectMember(projectMember)
	if err != nil {
		return errors.New("there was some problem deleting the member")
	}
	return nil
}
