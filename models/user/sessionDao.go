package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"slashbase.com/backend/config"
	"slashbase.com/backend/db"
)

func (d Dao) CreateUserSession(session *UserSession) error {
	result := db.GetDB().Create(session)
	return result.Error
}

func (d Dao) GetUserSessionByID(sessionID string) (*UserSession, error) {
	var userSession UserSession
	err := db.GetDB().Where(&UserSession{ID: sessionID}).Preload("User").First(&userSession).Error
	return &userSession, err
}

func (d Dao) GetUserSessionBySecret(secret string) (*UserSession, error) {
	var userSession UserSession
	err := db.GetDB().Where(&UserSession{Secret: sql.NullString{String: secret, Valid: true}}).Preload("User").First(&userSession).Error
	return &userSession, err
}

func (d Dao) GetUserSessionFromMagicLinkToken(tokenString string) (*UserSession, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetMagicLinkTokenSecret()), nil
	})
	if err != nil {
		return nil, errors.New("Invalid Token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionSecret := claims["sessionSecret"].(string)
		session, err := d.GetUserSessionBySecret(sessionSecret)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("Invalid Token")
			}
			return nil, errors.New("There was some problem")
		}
		return session, nil
	}
	return nil, errors.New("Invalid Token")
}

func (d Dao) GetUserSessionFromAuthToken(tokenString string) (*UserSession, error) {
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
