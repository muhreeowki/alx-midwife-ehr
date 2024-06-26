package database

import (
	"github.com/muhreeowki/midwifery-ehr/internal/models"
)

/* CreatePatient creates a new patient record in the database. */
func (engine *DatabaseEngine) CreatePatient(patientData *models.Patient) (patient models.Patient, err error) {
	patient = *patientData
	err = engine.DB.Create(patientData).Error
	if err != nil {
		return patient, err
	}
	return patient, nil
}

/* GetPatient retrieves a patient record from the database using the patient's id. */
func (engine *DatabaseEngine) GetPatient(id string) (patient models.Patient, err error) {
	err = engine.DB.First(&patient, id).Error
	if err != nil {
		return patient, err
	}
	return patient, nil
}

/* UpdatePatient updates a patient record in the database. */
func (engine *DatabaseEngine) UpdatePatient(patientData *models.Patient) (patient models.Patient, err error) {
	patient = *patientData
	err = engine.DB.Save(patient).Error
	if err != nil {
		return patient, err
	}
	return patient, nil
}

/* DeletePatient deletes a patient record from the database using the patient's id. */
func (engine *DatabaseEngine) DeletePatient(id string) (err error) {
	err = engine.DB.Delete(&models.Patient{}, id).Error
	return err
}
