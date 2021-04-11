package daos

import (
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
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

func (d UserDao) GetUserWithTeam(userID string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where(&models.User{ID: userID}).Preload("Teams").First(&user).Error
	return &user, err
}
