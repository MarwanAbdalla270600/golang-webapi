package auth

import (
	"carsharing/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginData user.UserLoginParser

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginService(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("session_id", token, int(user.SESSION_EXPIRE_TIME.Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func LogoutHandler(c *gin.Context) {
	sessionVal, _ := c.Get("session")
	session := sessionVal.(string)

	success := LogoutService(session)
	if !success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Something went wrong with logout"})
		return
	}

	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func RegisterHandler(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	RegisterService(&user)
	c.JSON(http.StatusOK, user)
}
