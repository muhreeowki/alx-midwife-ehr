package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	// Patient's personal details
	FirstName   string    `json:"firstName" gorm:"not null;"`
	LastName    string    `json:"lastName" gorm:"not null;"`
	BirthDate   time.Time `json:"birthDate"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	PartnerName string    `json:"partnerName"`
	ImageURL    string    `json:"imageURL"`
	// Patients' medical details
	LMP            time.Time `json:"lmp"`            // Last Menstrual Period
	ConceptionDate time.Time `json:"conceptionDate"` // Date of conception
	SonoDate       time.Time `json:"sonoDate"`       // Date of sonogram
	CRL            float64   `json:"crl"`            // Crown Rump Length
	CRLDate        time.Time `json:"crlDate"`        // Date of CRL
	EDD            time.Time `json:"edd"`            // Estimated Due Date
	RhFactor       string    `json:"rhFactor"`       // Rh Factor
	// Delivery details
	Delivered    bool      `json:"delivered" gorm:"default:false"` // Has the patient delivered
	DeliveryDate time.Time `json:"deliveryDate"`                   // Date of delivery
	// Midwife details
	MidwifeID uint32 `json:"midwifeID" gorm:"not null;"`
}
