package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
)

func getAuthUserHasAdminRoleForProject(authUser *models.User, projectID string) (bool, error) {
	pMember, err := getIfAuthUserProjectMemberForProject(authUser, projectID)
	if err != nil {
		return false, err
	}

	if pMember.Role.Name == models.ROLE_ADMIN {
		return true, nil
	}
	return false, errors.New("not allowed")
}

func getIfAuthUserProjectMemberForProject(authUser *models.User, projectID string) (*models.ProjectMember, error) {
	pMember, notFound, err := dao.ProjectMember.FindProjectMember(projectID, authUser.ID)
	if err != nil {
		if notFound {
			return nil, errors.New("not allowed")
		}
		return nil, errors.New("there was some problem")
	}

	return pMember, nil
}
