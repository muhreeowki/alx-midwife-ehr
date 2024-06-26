package database

import (
	"gorm.io/gorm"
)

func (engine *DatabaseEngine) CreatePatient(patient *Patient) (tx *gorm.DB) {
	tx = engine.DB.Create(patient)
	return tx
}

func (engine *DatabaseEngine) GetPatient(id string) (patient Patient, err error) {
	tx := engine.DB.First(&patient, id)
	if tx.Error != nil {
		return patient, tx.Error
	}
	return patient, nil
}