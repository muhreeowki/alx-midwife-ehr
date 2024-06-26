package database

import (
	"gorm.io/gorm"
)

/* CreatePatient creates a new patient record in the database. */
func (engine *DatabaseEngine) CreatePatient(patient *Patient) (tx *gorm.DB) {
	tx = engine.DB.Create(patient)
	return tx
}

/* GetPatient retrieves a patient record from the database using the patient's id. */
func (engine *DatabaseEngine) GetPatient(id string) (patient Patient, err error) {
	tx := engine.DB.First(&patient, id)
	if tx.Error != nil {
		return patient, tx.Error
	}
	return patient, nil
}

/* UpdatePatient updates a patient record in the database. */
func (engine *DatabaseEngine) UpdatePatient(patient *Patient) (tx *gorm.DB) {
	tx = engine.DB.Save(patient)
	return tx
}
