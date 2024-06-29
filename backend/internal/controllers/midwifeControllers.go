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
	"gorm.io/gorm"
)

/* MidwifeSignupController returns a gin.HandlerFunc that parses a midwife record from the request body and calls the CreateMidwife method on the database engine to create a new midwife record. */
func MidwifeSignupController(c *gin.Context) {
	// Parse the request body into a Midwife struct
	var midwife models.Midwife
	if err := c.ShouldBindJSON(&midwife); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if the midwife already exists in the database
	var midwifeFound models.Midwife
	database.ENGINE.DB.Where("email=?", midwife.Email).Find(&midwifeFound)
	if midwifeFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
		return
	}
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(midwife.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	midwife.Password = string(passwordHash)
	// Create the midwife record in the database
	err = database.ENGINE.CreateMidwife(&midwife)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Return the created patient record in the response
	c.JSON(http.StatusCreated, gin.H{"data": midwife})
}

/* MidwifeLoginController is a Gin controller that handles the login endpoint for midwives. */
func MidwifeLoginController(c *gin.Context) {
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

/* MidwifeProfileController is a Gin controller that returns the midwife profile of the currently authenticated midwife. */
func MidwifeProfileController(c *gin.Context) {
	midwife, ok := c.Get("currentMidwife")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email and password was found"})
		return
	}
	// Return the midwife in the response
	c.JSON(http.StatusOK, gin.H{"midwife": midwife})
}

/* MidwifePatientsController is a Gin controller that returns all patients associated with a midwife. */
func MidwifePatientsController(c *gin.Context) {
	// Get the current midwifeData from the context
	midwifeData, ok := c.Get("currentMidwife")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user with that email and password was found"})
		return
	}
	// Convert the midwifeData to a Midwife struct
	midwife := models.Midwife{
		Model: gorm.Model{ID: midwifeData.(models.AuthMidwifeOutput).ID},
		Name:  midwifeData.(models.AuthMidwifeOutput).Name,
		Email: midwifeData.(models.AuthMidwifeOutput).Email,
	}
	// Get all patients associated with the midwife
	patients, err := database.ENGINE.GetMidwifePatients(midwife)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Return the patients in the response
	c.JSON(http.StatusOK, gin.H{"patients": patients})
}
