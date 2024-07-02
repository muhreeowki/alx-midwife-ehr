package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/models"
)

/* CreatePatientController returns a gin.HandlerFunc that parses a patient record from the request body and calls the CreatePatient method on the database engine to create a new patient record. */
func CreatePatientController(c *gin.Context) {
	// Parse the request body into a Patient struct
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if the patient name is provided
	if patient.FirstName == "" && patient.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provide patient name"})
		return
	}
	// Set the midwife ID from the authenticated midwife if available
	midwifeId := c.MustGet("currentMidwife")
	patient.MidwifeID = uint32(midwifeId.(models.AuthMidwifeOutput).ID)
	// Create the patient record in the database
	err := database.ENGINE.CreatePatient(&patient)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the created patient record in the response
	c.JSON(http.StatusCreated, gin.H{"patient": patient})
}

/* GetPatientController returns a gin.HandlerFunc that parses a patient ID from the URL and calls the GetPatient method on the database engine to retrieve the patient record. */
func GetPatientController(c *gin.Context) {
	// Get the patient ID from the URL
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil || numId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}
	// Get the patient record from the database
	patient, err := database.ENGINE.GetPatient(id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// Check if the midwife is authorized to access the patient record
	midwifeId := c.MustGet("currentMidwife")
	if patient.MidwifeID != uint32(midwifeId.(models.AuthMidwifeOutput).ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Return the patient record in the response
	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

/* UpdatePatientController returns a gin.HandlerFunc that parses a patient record from the request body and calls the UpdatePatient method on the database engine to update the patient record. */
func UpdatePatientController(c *gin.Context) {
	// Parse the request body into a Patient struct
	var patientParams models.Patient
	if err := c.ShouldBindJSON(&patientParams); err != nil || patientParams.ID == 0 {
		if err != nil {
			c.Error(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	// Get the patient record from the database
	patient, err := database.ENGINE.GetPatient(strconv.Itoa(int(patientParams.ID)))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	// Check if the midwife is authorized to update the patient record
	midwifeId := c.MustGet("currentMidwife")
	if patient.MidwifeID != uint32(midwifeId.(models.AuthMidwifeOutput).ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Update the patient record in the database
	err = database.ENGINE.UpdatePatient(&patientParams)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the updated patient record from the database
	patient, err = database.ENGINE.GetPatient(strconv.Itoa(int(patientParams.ID)))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"patient": patient})
}

func DeletePatientController(c *gin.Context) {
	// Get the patient ID from the URL
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err != nil || numId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}
	// Get the patient record from the database
	dbPatient, err := database.ENGINE.GetPatient(id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	// Check if the midwife is authorized to update the patient record
	midwifeId := c.MustGet("currentMidwife")
	if dbPatient.MidwifeID != uint32(midwifeId.(models.AuthMidwifeOutput).ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Delete the patient record from the database
	err = database.ENGINE.DeletePatient(id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return a success message in the response
	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
