package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, GetAllUsersService())
}
