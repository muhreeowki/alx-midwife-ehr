package main

import (
	"fmt"
	"os"

	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/initializers"
	"github.com/muhreeowki/midwifery-ehr/internal/server"
)

/* Init initializes the application */
func init() {
	// Load environment variables
	err := initializers.LoadEnvVariables()
	if err != nil {
		panic("Failed to load environment variables")
	}

	// Connect to the database
	err = database.ConnectToDB()
	if err != nil {
		panic("Failed to connect to database")
	}
}

func main() {
	// Start up the server
	r := server.SetupRouter()
	// Listen and serve
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
