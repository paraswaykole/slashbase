package dao

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/slashbaseide/slashbase/internal/common/config"
	"github.com/slashbaseide/slashbase/internal/common/db"
	"github.com/slashbaseide/slashbase/internal/server/models"
	"gorm.io/gorm"
)

func (userDao) CreateUserSession(session *models.UserSession) error {
	result := db.GetDB().Create(session)
	return result.Error
}

func (userDao) GetUserSessionByID(sessionID string) (*models.UserSession, error) {
	var userSession models.UserSession
	err := db.GetDB().Where(&models.UserSession{ID: sessionID}).Preload("User.Projects").First(&userSession).Error
	return &userSession, err
}

func (d userDao) GetUserSessionFromAuthToken(tokenString string) (*models.UserSession, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfig().AuthTokenSecret), nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionID := claims["sessionID"].(string)
		session, err := d.GetUserSessionByID(sessionID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("invalid token")
			}
			return nil, errors.New("there was some problem")
		}
		if !session.IsActive {
			return nil, errors.New("invalid Token")
		}
		return session, nil
	}
	return nil, errors.New("invalid Token")
}
