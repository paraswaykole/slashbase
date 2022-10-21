package controllers

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/daos"
	"slashbase.com/backend/internal/models"
)

type UserController struct{}

var userDao daos.UserDao

func (uc UserController) LoginUser(email, password string) (*models.UserSession, error) {

	usr, err := userDao.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid user")
		}
		return nil, errors.New("there was some problem")
	}
	if usr.VerifyPassword(password) {
		userSession, _ := models.NewUserSession(usr.ID)
		err = userDao.CreateUserSession(userSession)
		userSession.User = *usr
		if err != nil {
			return nil, errors.New("there was some problem")
		}
		return userSession, nil
	}
	return nil, errors.New("invalid user")
}

func (uc UserController) EditAccount(authUser *models.User, name, profileImageUrl string) error {

	err := userDao.EditUser(authUser.ID, name, profileImageUrl)
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

func (uc UserController) ChangePassword(authUser *models.User, oldPassword, newPassword string) error {

	isOldPaswordValid := authUser.VerifyPassword(oldPassword)
	if !isOldPaswordValid {
		return errors.New("old password is incorrect")
	}

	err := authUser.SetPassword(newPassword)
	if err != nil {
		return errors.New("there was some problem")
	}

	err = userDao.UpdatePassword(authUser.ID, authUser.Password)
	if err != nil {
		return errors.New("there was some problem")
	}

	return nil
}

func (uc UserController) GetUsersPaginated(authUser *models.User, searchTerm string, offset int) (*[]models.User, int, error) {

	if !authUser.IsRoot {
		return nil, 0, errors.New("not allowed")
	}

	var users *[]models.User
	var err error
	if searchTerm == "" {
		users, err = userDao.GetUsersPaginated(offset)
	} else {
		users, err = userDao.SearchUsersPaginated(searchTerm, offset)
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

func (uc UserController) AddUser(authUser *models.User, email, password string) error {
	if !authUser.IsRoot {
		return errors.New("not allowed")
	}
	usr, err := userDao.GetUserByEmail(email)
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
	err = userDao.CreateUser(usr)
	if err != nil {
		return errors.New("there was some problem")
	}
	return nil
}
