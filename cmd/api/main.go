package main

import (
	"context"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"golang-apis-kickstart/internal/config"
	"golang-apis-kickstart/internal/database"
	"golang-apis-kickstart/internal/routes"
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

	// setup middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(helmet.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // change this to your domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Use(stats.RequestStats())

	routes.SetupRoutes(router)

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
