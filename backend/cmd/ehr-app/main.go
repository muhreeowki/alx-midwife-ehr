package main

import (
	"fmt"
	"os"

	"github.com/muhreeowki/midwifery-ehr/internal/database"
	"github.com/muhreeowki/midwifery-ehr/internal/initializers"
	"github.com/muhreeowki/midwifery-ehr/internal/server"
)

/* Init initializes the application */
func Init() {
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
	// Initialize the application
	Init()
	// Start up the server
	r := server.SetupRouter(database.ENGINE)
	// Listen and serve
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
