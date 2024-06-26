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
	// Connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schemas
	err = db.AutoMigrate(&Patient{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Midwife{})
	if err != nil {
		return err
	}

	ENGINE = &DatabaseEngine{DB: db}

	return nil
}
