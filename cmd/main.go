package main

import (
	"carsharing/database"
	"carsharing/internal/auth"
	"carsharing/internal/middleware"
	"carsharing/internal/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, handler *auth.AuthHandler) {
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", handler.Register)
		authRoutes.POST("/login", handler.Login)
		authRoutes.POST("/logout", middleware.AuthMiddleware(), handler.Logout)
	}
	userRoutes := router.Group("/api/users")
	{
		userRoutes.GET("/", middleware.AuthMiddleware(), user.GetAllHandler)
	}
}

func main() {
	db := database.ConnectDB()
	authHandler := &auth.AuthHandler{DB: db}
	err := database.RunSQLFile(db, "database/init.sql")
	if err != nil {
		fmt.Printf("can not load init sql: %s", err)
	} else {
		fmt.Println("Initialize Database")
	}
	router := gin.Default()
	initRoutes(router, authHandler)
	router.Run(":8080")
}
