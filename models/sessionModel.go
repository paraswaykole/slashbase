package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
)

type UserSession struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    string         `gorm:"not null"`
	Secret    sql.NullString `gorm:"index"`
	IsActive  bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignkey:user_id"`
}

func NewUserSession(userID string) (*UserSession, error) {
	var err error = nil
	if userID == "" {
		return nil, errors.New("user id cannot be empty")
	}
	newSecret := strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "")
	return &UserSession{
		UserID: userID,
		Secret: sql.NullString{
			String: newSecret,
			Valid:  true,
		},
		IsActive: false,
	}, err
}

func (session UserSession) GetMagicLink() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionSecret": session.Secret.String,
	})
	tokenString, err := token.SignedString(config.GetMagicLinkTokenSecret())
	if err != nil {
		panic(err)
	}
	return config.GetAppHost() + "/login?t=" + tokenString
}

func (session UserSession) GetAuthToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionID": session.ID,
	})
	tokenString, err := token.SignedString(config.GetAuthTokenSecret())
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (session UserSession) SetActive() error {
	session.IsActive = true
	session.Secret.Valid = false
	return session.Save()
}

func (session UserSession) SetInActive() error {
	session.IsActive = false
	session.Secret.Valid = false
	return session.Save()
}

func (session UserSession) Save() error {
	return db.GetDB().Save(&session).Error
}
