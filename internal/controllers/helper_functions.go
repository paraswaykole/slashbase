package controllers

import (
	"errors"

	"slashbase.com/backend/internal/dao"
	"slashbase.com/backend/internal/models"
	"slashbase.com/backend/pkg/queryengines/queryconfig"
)

func getAuthUserHasAdminRoleForProject(authUser *models.User, projectID string) (bool, error) {
	pMember, notFound, err := dao.Project.FindProjectMember(projectID, authUser.ID)
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
	pMember, notFound, err := dao.Project.FindProjectMember(projectID, authUser.ID)
	if err != nil {
		if notFound {
			return nil, errors.New("not allowed")
		}
		return nil, errors.New("there was some problem")
	}

	return pMember, nil
}

func getQueryConfigsForProjectMember(projectMember *models.ProjectMember, dbConn *models.DBConnection) *queryconfig.QueryConfig {
	createLog := func(query string) {
		queryLog := models.NewQueryLog(projectMember.UserID, dbConn.ID, query)
		go dao.DBQueryLog.CreateDBQueryLog(queryLog)
	}
	rolePermissions, _ := dao.RolePermission.GetRolePermissionsForRole(projectMember.RoleID)
	readOnly := false
	for _, perm := range *rolePermissions {
		if perm.Name == models.ROLE_PERMISSION_NAME_READ_ONLY {
			readOnly = true
		}
	}
	return queryconfig.NewQueryConfig(readOnly, createLog)
}
