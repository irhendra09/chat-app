package repository

import (
	"time"

	"github.com/irhendra09/chat-app/app/models"
	"github.com/irhendra09/chat-app/pkg/database"
)

func InsertNewUser(user *models.User) error {
	return database.DB.Create(&user).Error
}

func InsertNewUserSession(session *models.UserSession) error {
	return database.DB.Create(&session).Error
}

func GetUserSessionByToken(token string) (models.UserSession, error) {
	var (
		resp models.UserSession
		err  error
	)

	err = database.DB.Where("token = ?", token).Last(&resp).Error
	return resp, err
}

func DeleteUserSessionByToken(token string) error {
	return database.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func UpdateUserSessionToken(token string, tokenExpired time.Time, refreshToken string) error {
	return database.DB.Exec("UPDATE user_sessions SET token = ?, token_expired=? WHERE refresh_token = ?", token, tokenExpired, refreshToken).Error
}

func GetUserByUsername(username string) (models.User, error) {
	var (
		resp models.User
		err  error
	)
	err = database.DB.Where("username = ?", username).Last(&resp).Error
	return resp, err
}
