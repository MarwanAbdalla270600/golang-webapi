package user

import "time"

const SESSION_EXPIRE_TIME = 24 * time.Hour

type User struct {
	Username          string    `json:"username"`
	Password          string    `json:"-"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Session           string    `json:"token"`
	SessionExpireDate time.Time `json:"-"`
}

type UserLoginParser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
