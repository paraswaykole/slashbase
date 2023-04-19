package models

import (
	"database/sql"
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              string `gorm:"type:uuid;primaryKey"`
	Email           string `gorm:"unique;not null"`
	Password        string `gorm:"not null"`
	IsRoot          bool   `gorm:"not null;default:false"`
	FullName        sql.NullString
	ProfileImageURL sql.NullString
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	Projects []models.Project `gorm:"many2many:project_members;"`
}

func NewUser(email, textPassword string) (*User, error) {
	var err error = nil
	if email == "" || textPassword == "" {
		return nil, errors.New("fields cannot be empty")
	}
	re := regexp.MustCompile(`^\w+([\.\+-]?\w+)*@\w+([\.-]?\w+)*(\.\w+)+$`)
	if !re.Match([]byte(email)) {
		return nil, errors.New("email id is not valid")
	}
	user := User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: textPassword,
	}
	user.hashPassword()
	return &user, err
}

func (u *User) SetPassword(textPassword string) error {
	u.Password = textPassword
	return u.hashPassword()
}

func (u *User) VerifyPassword(textPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(textPassword))
	return err == nil
}

func (u *User) hashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(bytes)
	return err
}
