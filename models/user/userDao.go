package user

import (
	"slashbase.com/backend/db"
)

type Dao struct{}

func (d Dao) CreateUser(user *User) error {
	result := db.GetDB().Create(user)
	return result.Error
}

func (d Dao) GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.GetDB().Where(&User{Email: email}).First(&user).Error
	return &user, err
}
