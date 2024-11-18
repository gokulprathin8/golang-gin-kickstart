package routes

import (
	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
	"golang-apis-kickstart/internal/controllers"
	"net/http"
	"os"
)

func SetupRoutes(router *gin.Engine) {
	// load templates
	workingDir, _ := os.Getwd()
	templatePath := workingDir + "/internal/templates/**/*"
	router.LoadHTMLGlob(templatePath)

	statsReportRoutes(router)
	userRoutes(router)

	homePage(router)
}

func statsReportRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})
}

func userRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.CreateUser)
}

func homePage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/index.html", gin.H{
			"title": "Welcome to Golang APIs Kickstart",
		})
	})
}
