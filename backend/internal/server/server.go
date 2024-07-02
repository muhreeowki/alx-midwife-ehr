package server

import (
	"github.com/gin-gonic/gin"
	"github.com/muhreeowki/midwifery-ehr/internal/auth"
	"github.com/muhreeowki/midwifery-ehr/internal/controllers"
)

/* SetupRouter configures the Gin router with the necessary routes and middleware and returns the configured router. */
func SetupRouter() *gin.Engine {
	// Set the router as the default one provided by Gin
	r := gin.Default()

	// Setup test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Setup patient CRUD endpoints
	r.POST("/api/patient", auth.CheckAuth, controllers.CreatePatientController)
	r.GET("/api/patient/:id", auth.CheckAuth, controllers.GetPatientController)
	r.PATCH("/api/patient", auth.CheckAuth, controllers.UpdatePatientController)
	r.DELETE("/api/patient/:id", auth.CheckAuth, controllers.DeletePatientController)

	// Setup midwife CRUD endpoints
	r.POST("/api/auth/signup", controllers.MidwifeSignupController)
	r.POST("/api/auth/login", controllers.MidwifeLoginController)
	r.GET("/api/auth/profile", auth.CheckAuth, controllers.MidwifeProfileController)
	r.GET("api/mypatients", auth.CheckAuth, controllers.MidwifePatientsController)

	return r
}
