package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang-apis-kickstart/internal/config"
	"golang-apis-kickstart/internal/database"
	"golang-apis-kickstart/internal/dto"
	"golang-apis-kickstart/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func CreateUser(c *gin.Context) {
	var authInput dto.AuthInput

	if err := c.ShouldBind(&authInput); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		c.Abort()
		return
	}

	// if a user with the same email already exists, return an error
	var userFound models.User
	database.DB.Where("email = ?", authInput.Email).First(&userFound)

	if len(userFound.Email) > 0 {
		c.JSON(400, gin.H{
			"error": "User with this email already exists",
		})
		c.Abort()
		return
	}

	// generate a hash for the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate password hash",
		})
	}

	user := models.User{
		Email:    authInput.Email,
		Password: string(hashedPassword),
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	var authInput dto.AuthInput

	if err := c.ShouldBind(&authInput); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		c.Abort()
		return
	}

	var user models.User
	database.DB.Where("email = ?", authInput.Email).First(&user)

	if len(user.Email) == 0 {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		c.Abort()
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authInput.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		c.Abort()
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(config.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
