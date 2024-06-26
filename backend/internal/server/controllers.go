package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
)

/* CreatePatientController creates a new patient record in the database. */
func CreatePatientController(engine *database.DatabaseEngine) func(*gin.Context) {
	return func(c *gin.Context) {
		var patient database.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tx := engine.CreatePatient(&patient)
		if err := tx.Error; err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Patient created successfully", "data": tx})
	}
}
