package auth

import (
	"carsharing/internal/user"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Register(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	RegisterService(&u, h) // or use h.DB if DB access is added
	c.JSON(http.StatusOK, u)
}

func (h *AuthHandler) Login(c *gin.Context) {
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

func (h *AuthHandler) Logout(c *gin.Context) {
	sessionVal, _ := c.Get("session")
	session := sessionVal.(string)

	if !LogoutService(session) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Something went wrong with logout"})
		return
	}

	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
