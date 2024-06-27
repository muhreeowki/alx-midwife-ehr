package auth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/models"
)

func CheckAuth(c *gin.Context) {
	// ********** AUTHORIZATION HEADER **********
	// Check if the Authorization header is present
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check if the Authorization header is in the correct format
	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is in the wrong format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ********** JWT AUTHENTICATION **********
	tokenString := authToken[1]
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	// Check if the token is valid
	if err != nil || !token.Valid {
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Get the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check if the token has expired
	if float64(time.Now().Unix()) > claims["expires_at"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token has expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Get the midwife id from the token claims and check if the midwife exists
	midwife, err := database.ENGINE.GetMidwife(claims["midwife_id"].(float64))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email and password was found"})
		return
	}
	// Set the midwife in the context
	c.Set("currentMidwife", models.AuthMidwifeOutput{
		ID:    midwife.ID,
		Email: midwife.Email,
		Name:  midwife.Name,
	})
	c.Next()
}
