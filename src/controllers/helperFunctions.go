package controllers

import (
	"errors"

	"slashbase.com/backend/src/models"
	"slashbase.com/backend/src/utils"
)

func GetAuthUserHasRolesForProject(authUser *models.User, projectID string, hasRoles []string) (bool, error) {
	pMember, notFound, err := projectDao.FindProjectMember(projectID, authUser.ID)
	if notFound {
		return false, errors.New("not allowed")
	} else if err != nil {
		return false, errors.New("there was some problem")
	}
	if utils.ContainsString(hasRoles, pMember.Role) {
		return true, nil
	}
	return false, errors.New("not allowed")
}

func GetAuthUserProjectMemberForProject(authUser *models.User, projectID string) (*models.ProjectMember, error) {
	pMember, notFound, err := projectDao.FindProjectMember(projectID, authUser.ID)
	if err != nil {
		if notFound {
			return nil, errors.New("not allowed")
		}
		return nil, errors.New("there was some problem")
	}

	return pMember, nil
}
