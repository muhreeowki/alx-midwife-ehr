package server

import (
	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/controllers"
	"github.com/muhreeowki/midwifery-ehr/internal/database"
)

/* SetupRouter configures the Gin router with the necessary routes and middleware and returns the configured router. */
func SetupRouter(engine *database.DatabaseEngine) *gin.Engine {
	// Set the router as the default one provided by Gin
	r := gin.Default()

	// Setup test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Setup patient CRUD endpoints
	r.POST("/patient", controllers.CreatePatientController(engine))
	r.GET("/patient/:id", controllers.GetPatientController(engine))
	r.PATCH("/patient", controllers.UpdatePatientController(engine))
	r.DELETE("/patient/:id", controllers.DeletePatientController(engine))

	return r
}
