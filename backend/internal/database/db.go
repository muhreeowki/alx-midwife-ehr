package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseEngine struct {
	DB *gorm.DB
}

var ENGINE *DatabaseEngine

/* ConnectToDB creates a connection to the database, migrates the schema and returns a gorm.DB object or an error if it fails. */
func ConnectToDB() error {
	// ********** Connect to the database **********
	// Connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)
	// Setup the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// TODO: Consider moving the migration to a separate function
	// ********** Migrate the schema **********
	// Patients table
	err = db.AutoMigrate(&Patient{})
	if err != nil {
		return err
	}
	// Midwives table
	err = db.AutoMigrate(&Midwife{})
	if err != nil {
		return err
	}

	// ********** Set the global ENGINE variable **********
	ENGINE = &DatabaseEngine{DB: db}
	return nil
}
