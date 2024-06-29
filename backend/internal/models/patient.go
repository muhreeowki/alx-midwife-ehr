package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	// Patient's personal details
	FirstName   string         `json:"first_name" gorm:"not null;"`
	LastName    string         `json:"last_name" gorm:"not null;"`
	BirthDate   sql.NullTime   `json:"birth_date"`
	Email       sql.NullString `json:"email"`
	Phone       sql.NullString `json:"phone"`
	Address     sql.NullString `json:"address"`
	PartnerName sql.NullString `json:"partner_name"`
	ImageURL    sql.NullString `json:"image_url"`

	// Patients' medical details
	LMP            sql.NullTime    `json:"lmp"`             // Last Menstrual Period
	ConceptionDate sql.NullTime    `json:"conception_date"` // Date of conception
	SonoDate       sql.NullTime    `json:"sono_date"`       // Date of sonogram
	CRL            sql.NullFloat64 `json:"crl"`             // Crown Rump Length
	CRLDate        sql.NullTime    `json:"crl_date"`        // Date of CRL
	EDD            sql.NullTime    `json:"edd"`             // Estimated Due Date
	RhFactor       sql.NullString  `json:"rh_factor"`       // Rh Factor

	// Delivery details
	Delivered    bool         `json:"delivered" gorm:"default:false"` // Has the patient delivered
	DeliveryDate sql.NullTime `json:"delivery_date"`                  // Date of delivery

	// Midwife details
	MidwifeID uint32 `json:"midwife_id" gorm:"not null;"`
}
