package auth

import (
	userpkg "carsharing/internal/user"
	"carsharing/internal/utils"
	"errors"
	"sync"
	"time"
)

var SessionMap sync.Map

func GetUsernameFromSessionMapService(session string) (string, bool) {
	usernameAny, ok := SessionMap.Load(session)
	username, _ := usernameAny.(string) // ✅ type assertion
	return username, ok
}

func LoginService(username string, password string) (string, error) {
	user, ok := userpkg.UserMap[username]
	if !ok || user.Password != password {
		return "", errors.New("username does not exist or invalid password")
	}
	session, _ := utils.GenerateRandomToken(32)
	user.Session = session
	user.SessionExpireDate = time.Now().Add(userpkg.SESSION_EXPIRE_TIME)
	SessionMap.Store(session, username)
	return user.Session, nil
}

func LogoutService(sessionId string) bool {
	usernameAny, ok := SessionMap.Load(sessionId)
	if !ok {
		return false
	}

	username, ok := usernameAny.(string)
	if !ok {
		return false
	}

	user, ok := userpkg.GetUserByUsernameService(username)
	if !ok || user == nil {
		return false
	}

	user.Session = ""
	user.SessionExpireDate = time.Time{}
	SessionMap.Delete(sessionId)
	return true
}

func RegisterService(user *userpkg.User) {
	userpkg.UserMap[user.Username] = user

}
