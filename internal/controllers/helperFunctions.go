package controllers

import (
	"errors"

	"slashbase.com/backend/internal/models"
)

func getAuthUserHasAdminRoleForProject(authUser *models.User, projectID string) (bool, error) {
	pMember, notFound, err := projectDao.FindProjectMember(projectID, authUser.ID)
	if notFound {
		return false, errors.New("not allowed")
	} else if err != nil {
		return false, errors.New("there was some problem")
	}
	if pMember.Role.Name == models.ROLE_ADMIN {
		return true, nil
	}
	return false, errors.New("not allowed")
}

func getAuthUserProjectMemberForProject(authUser *models.User, projectID string) (*models.ProjectMember, error) {
	pMember, notFound, err := projectDao.FindProjectMember(projectID, authUser.ID)
	if err != nil {
		if notFound {
			return nil, errors.New("not allowed")
		}
		return nil, errors.New("there was some problem")
	}

	return pMember, nil
}
