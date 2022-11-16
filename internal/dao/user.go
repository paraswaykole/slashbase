package dao

import (
	"database/sql"
	"strings"

	"slashbase.com/backend/internal/config"
	"slashbase.com/backend/internal/db"
	"slashbase.com/backend/internal/models"
)

type userDao struct{}

var User userDao

func (userDao) CreateUser(user *models.User) error {
	result := db.GetDB().Create(user)
	return result.Error
}

func (userDao) GetRootUserOrCreate(user models.User) (*models.User, error) {
	var result models.User
	err := db.GetDB().Model(&models.User{}).Where(&models.User{Email: user.Email}).Attrs(&user).FirstOrCreate(&result).Error
	return &result, err
}

func (userDao) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where(&models.User{Email: email}).First(&user).Error
	return &user, err
}

func (userDao) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where(&models.User{ID: userID}).Preload("Projects").First(&user).Error
	return &user, err
}

func (userDao) EditUser(userID string, name string, profileImageURL string) error {
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

func (userDao) UpdatePassword(userID string, newPasswordHash string) error {
	err := db.GetDB().Where(&models.User{ID: userID}).Updates(&models.User{
		Password: newPasswordHash,
	}).Error
	return err
}

func (userDao) GetUsersPaginated(offset int) (*[]models.User, error) {
	var users []models.User
	err := db.GetDB().
		Model(&models.User{}).
		Offset(offset).Limit(config.PAGINATION_COUNT).
		Preload("Projects").Find(&users).Error
	return &users, err
}

func (userDao) SearchUsersPaginated(searchTerm string, offset int) (*[]models.User, error) {
	searchTerm = "%" + strings.ToLower(searchTerm) + "%"
	var users []models.User
	err := db.GetDB().
		Model(&models.User{}).
		Where("lower(email) LIKE ? OR lower(full_name) LIKE ?", searchTerm, searchTerm).
		Offset(offset).Limit(config.PAGINATION_COUNT).
		Preload("Projects").Find(&users).Error
	return &users, err
}
