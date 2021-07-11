package daos

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
	"slashbase.com/backend/models"
)

func (d UserDao) CreateUserSession(session *models.UserSession) error {
	result := db.GetDB().Create(session)
	return result.Error
}

func (d UserDao) GetUserSessionByID(sessionID string) (*models.UserSession, error) {
	var userSession models.UserSession
	err := db.GetDB().Where(&models.UserSession{ID: sessionID}).Preload("User.Teams").First(&userSession).Error
	return &userSession, err
}

func (d UserDao) GetUserSessionFromAuthToken(tokenString string) (*models.UserSession, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetAuthTokenSecret()), nil
	})
	if err != nil {
		return nil, errors.New("Invalid Token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionID := claims["sessionID"].(string)
		session, err := d.GetUserSessionByID(sessionID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("Invalid Token")
			}
			return nil, errors.New("There was some problem")
		}
		if !session.IsActive {
			return nil, errors.New("Invalid Token")
		}
		return session, nil
	}
	return nil, errors.New("Invalid Token")
}
