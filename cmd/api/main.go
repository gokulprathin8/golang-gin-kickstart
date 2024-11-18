package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-apis-kickstart/internal/config"
	"golang-apis-kickstart/internal/controllers"
	"golang-apis-kickstart/internal/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	srv := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Server gracefully shutdown")
	}
}
