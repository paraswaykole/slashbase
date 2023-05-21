package controllers

import (
	"errors"

	common "github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
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

func getQueryConfigsForProjectMember(projectMember *models.ProjectMember, dbConn *common.DBConnection) *qemodels.QueryConfig {
	createLog := func(query string) {
		queryLog := common.NewQueryLog(dbConn.ID, query)
		userQueryLog := models.NewUserQueryLog(*queryLog, projectMember.UserID)
		go dao.DBQueryLog.CreateDBQueryLog(userQueryLog)
	}
	rolePermissions, _ := dao.RolePermission.GetRolePermissionsForRole(projectMember.RoleID)
	readOnly := false
	for _, perm := range *rolePermissions {
		if perm.Name == models.ROLE_PERMISSION_NAME_READ_ONLY && perm.Value {
			readOnly = true
		}
	}
	return qemodels.NewQueryConfig(readOnly, createLog)
}
