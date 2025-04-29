package user

var UserMap = map[string]*User{ //todo, replace with database
	"marwan2000": {
		Username:  "marwan2000",
		Password:  "12345678",
		Firstname: "Marwan",
		Lastname:  "Abdalla",
	},
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
