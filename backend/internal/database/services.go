package database

import (
	"gorm.io/gorm"
)

/* CreatePatient creates a new patient record in the database. */
func (engine *DatabaseEngine) CreatePatient(patientData *Patient) (patient Patient, err error) {
	patient = *patientData
	tx := engine.DB.Create(patientData)
	if tx.Error != nil {
		return patient, tx.Error
	}
	return patient, nil
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
func (engine *DatabaseEngine) UpdatePatient(patientData *Patient) (patient Patient, err error) {
	patient = *patientData
	tx := engine.DB.Save(patient)
	if tx.Error != nil {
		return patient, tx.Error
	}
	return patient, nil
}

func (engine *DatabaseEngine) DeletePatient(id string) (tx *gorm.DB) {
	tx = engine.DB.Delete(&Patient{}, id)
	return tx
}
