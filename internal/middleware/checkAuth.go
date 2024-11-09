package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang-apis-kickstart/internal/config"
	"net/http"
	"strings"
	"time"
)

// Replace this with your actual secret key
var jwtSecret = []byte(config.SecretKey)

// CheckAuth middleware for validating JWT token
func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is present
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Missing authorization header",
		})
		c.Abort()
		return
	}

	// Split the Authorization header value
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization header format",
		})
		c.Abort()
		return
	}

	// Extract the token string
	tokenString := authToken[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, parseToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed to parse token",
		})
		c.Abort()
		return
	}

	// Check if the token is valid
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.Abort()
		return
	}

	// Extract claims and verify expiration
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		c.Abort()
		return
	}

	// Verify token expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token has expired",
			})
			c.Abort()
			return
		}
	}

	// Store claims in context for further use
	c.Set("claims", claims)

	// Proceed to the next handler
	c.Next()
}

// parseToken function for validating the signing method and returning the key
func parseToken(token *jwt.Token) (interface{}, error) {
	// Verify the signing method
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	// Return the secret key for validation
	return jwtSecret, nil
}
