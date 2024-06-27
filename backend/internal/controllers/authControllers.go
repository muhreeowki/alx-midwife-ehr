package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c *gin.Context) {
	// Parse the request body into a Login struct
	var loginInput models.AuthMidwifeInput
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the midwife exists in the database
	var midwife models.Midwife
	database.ENGINE.DB.Where("email=?", loginInput.Email).Find(&midwife)
	if midwife.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email and password was found"})
		return
	}

	// Compare the password hash
	err := bcrypt.CompareHashAndPassword([]byte(midwife.Password), []byte(loginInput.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email and password was found"})
		return
	}

	// Generate a JWT token
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"midwife_id": midwife.ID,
		"expires_at": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the user record in the response
	c.JSON(http.StatusOK, gin.H{"token": token})
}
