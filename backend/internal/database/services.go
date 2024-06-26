package database

import (
	"gorm.io/gorm"
)

func (engine *DatabaseEngine) CreatePatient(patient *Patient) (tx *gorm.DB) {
	tx = engine.DB.Create(patient)
	return tx
}
