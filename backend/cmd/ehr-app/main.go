package main

import (
	"fmt"

	"github.com/muhreeowki/midwifery-ehr/internal/database"
)

func main() {
	fmt.Println("Hello, World!")

	db, err := database.StartDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Successfully connected to database")

	midwife := database.Midwife{Email: "Midwife@gmail.com"}
	db.Find(midwife)
	fmt.Println(midwife)
}
