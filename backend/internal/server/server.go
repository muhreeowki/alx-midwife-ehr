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
	r.POST("/patient", auth.CheckAuth, controllers.CreatePatientController)
	r.GET("/patient/:id", auth.CheckAuth, controllers.GetPatientController)
	r.PATCH("/patient", auth.CheckAuth, controllers.UpdatePatientController)
	r.DELETE("/patient/:id", auth.CheckAuth, controllers.DeletePatientController)

	// Setup midwife CRUD endpoints
	r.POST("/signup", controllers.MidwifeSignupController)
	r.POST("/login", controllers.MidwifeLoginController)
	r.GET("/profile", auth.CheckAuth, controllers.GetMidwifeProfile)

	return r
}
