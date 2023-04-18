package controllers

import (
	"database/sql"
	"errors"

	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"gorm.io/gorm"
)

type UserController struct{}

func (UserController) LoginUser(email, password string) (*models.UserSession, error) {

	usr, err := dao.User.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid user")
		}
		return nil, errors.New("there was some problem")
	}
	if usr.VerifyPassword(password) {
		userSession, _ := models.NewUserSession(usr.ID)
		err = dao.User.CreateUserSession(userSession)
		userSession.User = *usr
		if err != nil {
			return nil, errors.New("there was some problem")
		}
		return userSession, nil
	}
	return nil, errors.New("invalid user")
}

func (UserController) EditAccount(authUser *models.User, name, profileImageUrl string) error {

	err := dao.User.EditUser(authUser.ID, name, profileImageUrl)
	if err != nil {
		return errors.New("there was some problem")
	}
	authUser.FullName = sql.NullString{
		String: name,
		Valid:  name != "",
	}
	authUser.ProfileImageURL = sql.NullString{
		String: profileImageUrl,
		Valid:  profileImageUrl != "",
	}

	return nil
}

func (UserController) ChangePassword(authUser *models.User, oldPassword, newPassword string) error {

	isOldPaswordValid := authUser.VerifyPassword(oldPassword)
	if !isOldPaswordValid {
		return errors.New("old password is incorrect")
	}

	err := authUser.SetPassword(newPassword)
	if err != nil {
		return errors.New("there was some problem")
	}

	err = dao.User.UpdatePassword(authUser.ID, authUser.Password)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}

func (UserController) GetUsersPaginated(authUser *models.User, searchTerm string, offset int) (*[]models.User, int, error) {

	if !authUser.IsRoot {
		return nil, 0, errors.New("not allowed")
	}

	var users *[]models.User
	var err error
	if searchTerm == "" {
		users, err = dao.User.GetUsersPaginated(offset)
	} else {
		users, err = dao.User.SearchUsersPaginated(searchTerm, offset)
	}

	if err != nil {
		return nil, 0, errors.New("there was some problem")
	}

	next := -1
	if len(*users) == config.PAGINATION_COUNT {
		next = offset + config.PAGINATION_COUNT
	}

	return users, next, nil
}

func (UserController) AddUser(authUser *models.User, email, password string) error {
	if !authUser.IsRoot {
		return errors.New("not allowed")
	}
	usr, err := dao.User.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			usr, err = models.NewUser(email, password)
			if err != nil {
				return err
			}
		} else {
			return errors.New("there was some problem")
		}
	}
	err = dao.User.CreateUser(usr)
	if err != nil {
		return errors.New("there was some problem")
	}
	return nil
}
