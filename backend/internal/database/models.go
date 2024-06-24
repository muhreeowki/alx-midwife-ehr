package database

import (
	"time"

	"gorm.io/gorm"
)

type Midwife struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
	Patients     []Patient
}

type Patient struct {
	gorm.Model
	// Patient's personal details
	FirstName   string
	LastName    string
	BirthDate   time.Time
	Email       string
	Phone       string
	Address     string
	PartnerName string

	// Patients' medical details
	LMP            time.Time // Last Menstrual Period
	ConceptionDate time.Time // Date of conception
	SonoDate       time.Time // Date of sonogram
	CRL            float64   // Crown Rump Length
	CRLDate        time.Time // Date of CRL measurement
	EDD            time.Time // Estimated Due Date
	RhFactor       string    // Rh factor

	// Delivery details
	Delivered    bool      // Has the patient delivered?
	DeliveryDate time.Time // Date of delivery

	// Midwife details
	MidwifeID uint32
}
