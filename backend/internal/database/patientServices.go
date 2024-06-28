package database

import (
	"github.com/muhreeowki/midwifery-ehr/internal/models"
)

/* CreatePatient creates a new patient record in the database. */
func (engine *DatabaseEngine) CreatePatient(patient *models.Patient) (err error) {
	err = engine.DB.Create(patient).Error
	return err
}

/* GetPatient retrieves a patient record from the database using the patient's id. */
func (engine *DatabaseEngine) GetPatient(id string) (patient models.Patient, err error) {
	err = engine.DB.First(&patient, id).Error
	return patient, err
}

// TODO: patient object prepopulates all fields, even if they are not provided in the request body. find a way to only update fields that are provided in the request body
/* UpdatePatient updates a patient record in the database. */
func (engine *DatabaseEngine) UpdatePatient(patient *models.Patient) (err error) {
	err = engine.DB.Model(patient).Updates(patient).Error
	return err
}

/* DeletePatient deletes a patient record from the database using the patient's id. */
func (engine *DatabaseEngine) DeletePatient(id string) (err error) {
	err = engine.DB.Delete(&models.Patient{}, id).Error
	return err
}
