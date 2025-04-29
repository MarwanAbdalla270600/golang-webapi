package user

import (
	"sync"
)

var UserMap = map[string]*User{
	"marwan2000": {
		Username:  "marwan2000",
		Password:  "12345678",
		Firstname: "Marwan",
		Lastname:  "Abdalla",
	},
}

var SessionMap sync.Map

func GetUsernameFromSessionMapService(session string) (string, bool) {
	usernameAny, ok := SessionMap.Load(session)
	username, _ := usernameAny.(string) // ✅ type assertion
	return username, ok
}

func GetUserByUsernameService(username string) (*User, bool) {
	user, ok := UserMap[username]
	if !ok {
		return nil, false
	}
	return user, true
}

func GetAllUsersService() []*User {
	var userList []*User

	for _, value := range UserMap {
		userList = append(userList, value)
	}
	return userList
}
