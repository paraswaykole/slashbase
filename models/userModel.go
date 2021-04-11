package models

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID              string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email           string `gorm:"unique;not null"`
	FullName        sql.NullString
	ProfileImageURL sql.NullString
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`

	Teams []Team `gorm:"many2many:team_members;"`
}

func NewUser(email string) (*User, error) {
	var err error = nil
	if email == "" {
		return nil, errors.New("Fields cannot be empty")
	}
	re := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w+)+$`)
	if !re.Match([]byte(email)) {
		return nil, errors.New("Email id is not valid")
	}
	return &User{
		Email: email,
	}, err
}
