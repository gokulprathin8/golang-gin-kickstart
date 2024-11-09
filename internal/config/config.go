package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var (
	ServerPort string
	SecretKey  string
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		gin.DefaultErrorWriter.Write([]byte("Failed to load .env file\n"))
		return
	}

	SecretKey = getEnv("SECRET_KEY", "")
	ServerPort = getEnv("SERVER_PORT", "8000")
	gin.DefaultWriter.Write([]byte(".env file loaded successfully\n"))

}

func getEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}
