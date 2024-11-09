package main

import (
	"github.com/gin-gonic/gin"
	"golang-apis-kickstart/internal/config"
	"golang-apis-kickstart/internal/controllers"
	"golang-apis-kickstart/internal/database"
)

func init() {
	config.LoadConfig()
	database.Init()
}

func main() {
	router := gin.Default()

	// create a group /v1 and add another group /auth in it
	v1 := router.Group("/v1")
	auth := v1.Group("/auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.CreateUser)
	//auth.POST("/logout", Logout)

	//v1.GET("/users", GetUsers)
	//v1.POST("/users", CreateUser)
	//v1.PUT("/users/:id", UpdateUser)
	//v1.DELETE("/users/:id", DeleteUser)

	err := router.Run(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
}
