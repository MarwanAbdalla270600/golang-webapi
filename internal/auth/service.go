package auth

import (
	userpkg "carsharing/internal/user"
	"carsharing/internal/utils"
	"errors"
	"time"
)

func LoginService(username string, password string) (string, error) {
	user, ok := userpkg.UserMap[username]
	if !ok || user.Password != password {
		return "", errors.New("username does not exist or invalid password")
	}
	session, _ := utils.GenerateRandomToken(32)
	user.Session = session
	user.SessionExpireDate = time.Now().Add(userpkg.SESSION_EXPIRE_TIME)
	userpkg.SessionMap.Store(session, username)
	return user.Session, nil
}

func RegisterService(user *userpkg.User) {
	userpkg.UserMap[user.Username] = user
}
