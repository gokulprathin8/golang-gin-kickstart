package main

import (
	"golang-apis-kickstart/internal/config"
	"golang-apis-kickstart/internal/database"
	"golang-apis-kickstart/internal/models"
)

func init() {

}

func init() {
	config.LoadConfig()
	database.Init()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
