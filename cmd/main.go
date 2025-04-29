package main

import (
	"carsharing/internal/auth"
	"carsharing/internal/middleware"
	"carsharing/internal/user"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", auth.RegisterHandler)
		authRoutes.POST("/login", auth.LoginHandler)
	}
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/", middleware.AuthMiddleware(), user.GetAllHandler)
	}
}

func main() {
	router := gin.Default()
	initRoutes(router)
	router.Run(":8080")
}
