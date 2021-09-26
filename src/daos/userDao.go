package daos

import (
	"database/sql"

	"slashbase.com/backend/src/db"
	"slashbase.com/backend/src/models"
)

type UserDao struct{}

func (d UserDao) CreateUser(user *models.User) error {
	result := db.GetDB().Create(user)
	return result.Error
}

func (d UserDao) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where(&models.User{Email: email}).First(&user).Error
	return &user, err
}

func (d UserDao) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where(&models.User{ID: userID}).Preload("Projects").First(&user).Error
	return &user, err
}

func (d UserDao) EditUser(userID string, name string, profileImageURL string) error {
	err := db.GetDB().Where(&models.User{ID: userID}).Updates(&models.User{
		FullName: sql.NullString{
			String: name,
			Valid:  name != "",
		},
		ProfileImageURL: sql.NullString{
			String: profileImageURL,
			Valid:  profileImageURL != "",
		},
	}).Error
	return err
}
