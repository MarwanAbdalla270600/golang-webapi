package middleware

import (
	"carsharing/internal/auth"
	"carsharing/internal/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing session_id"})
			c.Abort()
			return
		}
		username, ok := auth.GetUsernameFromSessionMapService(session)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			c.Abort()
			return
		}
		userModel, ok := user.GetUserByUsernameService(username)
		if !ok || userModel == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		if time.Now().After(userModel.SessionExpireDate) {
			userModel.Session = ""
			userModel.SessionExpireDate = time.Time{}
			auth.SessionMap.Delete(session)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			c.Abort()
			return
		}

		// If everything OK -> Save user info in context for the next handlers
		c.Set("user", userModel)
		c.Set("session", session)
		c.Next() // continue to the next handler
	}
}
