package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
)

/* CreatePatientController returns a gin.HandlerFunc that parses a patient record from the request body and calls the CreatePatient method on the database engine to create a new patient record. */
func CreatePatientController(engine *database.DatabaseEngine) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse the request body into a Patient struct
		var patient database.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// Create the patient record in the database
		tx := engine.CreatePatient(&patient)
		if err := tx.Error; err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		// Return the created patient record in the response
		c.JSON(http.StatusCreated, gin.H{"data": patient})
	}
}

/* GetPatientController returns a gin.HandlerFunc that parses a patient ID from the URL and calls the GetPatient method on the database engine to retrieve the patient record. */
func GetPatientController(engine *database.DatabaseEngine) func(*gin.Context) {
	return func(c *gin.Context) {
		// Get the patient ID from the URL
		id := c.Param("id")
		// Get the patient record from the database
		patient, err := engine.GetPatient(id)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		// Return the patient record in the response
		c.JSON(http.StatusOK, gin.H{"data": patient})
	}
}

/* UpdatePatientController returns a gin.HandlerFunc that parses a patient record from the request body and calls the UpdatePatient method on the database engine to update the patient record. */
func UpdatePatientController(engine *database.DatabaseEngine) func(*gin.Context) {
	return func(c *gin.Context) {
		var patient database.Patient
		if err := c.ShouldBindJSON(&patient); err != nil || patient.ID == 0 {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tx := engine.UpdatePatient(&patient)
		if err := tx.Error; err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": patient})
	}
}
