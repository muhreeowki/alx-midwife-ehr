package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/models"
	"golang.org/x/crypto/bcrypt"
)

/* CreateMidwifeController returns a gin.HandlerFunc that parses a midwife record from the request body and calls the CreateMidwife method on the database engine to create a new midwife record. */
func CreateMidwifeController(c *gin.Context) {
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

	// Create the patient record in the database
	err = database.ENGINE.CreateMidwife(&midwife)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Return the created patient record in the response
	c.JSON(http.StatusCreated, gin.H{"data": midwife})
}
