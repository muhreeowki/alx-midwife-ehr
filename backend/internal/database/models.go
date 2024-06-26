package database

import (
	"time"

	"gorm.io/gorm"
)

type Midwife struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null;" binding:"required"`
	Email        string `json:"email" gorm:"not null;unique" validate:"email" binding:"required"`
	PasswordHash string `json:"password" gorm:"not null;" binding:"required"`
	Patients     []Patient
}

type Patient struct {
	gorm.Model
	// Patient's personal details
	FirstName   string    `json:"first_name" gorm:"not null;" binding:"required"`
	LastName    string    `json:"last_name" gorm:"not null;" binding:"required"`
	BirthDate   time.Time `json:"birth_date"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	PartnerName string    `json:"partner_name"`

	// Patients' medical details
	LMP            time.Time `json:"lmp"`             // Last Menstrual Period
	ConceptionDate time.Time `json:"conception_date"` // Date of conception
	SonoDate       time.Time `json:"sono_date"`       // Date of sonogram
	CRL            float64   `json:"crl"`             // Crown Rump Length
	CRLDate        time.Time `json:"crl_date"`        // Date of CRL
	EDD            time.Time `json:"edd"`             // Estimated Due Date
	RhFactor       string    `json:"rh_factor"`       // Rh Factor

	// Delivery details
	Delivered    bool      `json:"delivered"`     // Has the patient delivered?
	DeliveryDate time.Time `json:"delivery_date"` // Date of delivery

	// Midwife details
	MidwifeID uint32 `json:"midwife_id" gorm:"not null;" binding:"required"`
}
