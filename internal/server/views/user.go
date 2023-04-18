package views

import (
	"time"

	"github.com/slashbaseide/slashbase/internal/server/models"
)

type UserView struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Name            *string   `json:"name"`
	IsRoot          bool      `json:"isRoot"`
	ProfileImageURL string    `json:"profileImageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type UserSessionView struct {
	ID        string    `json:"id"`
	User      UserView  `json:"user"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func BuildUser(usr *models.User) UserView {
	userView := UserView{
		ID:              usr.ID,
		Name:            nil,
		Email:           usr.Email,
		ProfileImageURL: usr.ProfileImageURL.String,
		IsRoot:          usr.IsRoot,
		CreatedAt:       usr.CreatedAt,
		UpdatedAt:       usr.UpdatedAt,
	}
	if usr.FullName.Valid {
		name := usr.FullName.String
		userView.Name = &name
	}
	return userView
}

func BuildUserSession(userSession *models.UserSession) UserSessionView {
	userSessView := UserSessionView{
		ID:        userSession.ID,
		User:      BuildUser(&userSession.User),
		IsActive:  userSession.IsActive,
		CreatedAt: userSession.CreatedAt,
		UpdatedAt: userSession.UpdatedAt,
	}
	return userSessView
}
