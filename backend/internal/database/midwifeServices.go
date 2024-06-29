package database

import "github.com/muhreeowki/midwifery-ehr/internal/models"

/* CreateMidwife creates a new midwife record in the database. */
func (engine *DatabaseEngine) CreateMidwife(midwife *models.Midwife) (err error) {
	err = engine.DB.Create(midwife).Error
	return err
}

/* GetMidwife retrieves a midwife record from the database. */
func (engine *DatabaseEngine) GetMidwife(id float64) (midwife models.Midwife, err error) {
	err = engine.DB.Where("id = ?", id).First(&midwife).Error
	return midwife, err
}

/* GetMidwifePatients retrieves all patients associated with a midwife. */
func (engine *DatabaseEngine) GetMidwifePatients(midwife models.Midwife) (patients []models.Patient, err error) {
	var populatedMidwife models.Midwife
	err = engine.DB.Model(&midwife).Where("id = ?", midwife.ID).Preload("Patients").Find(&populatedMidwife).Error
	patients = populatedMidwife.Patients
	return patients, err
}

// TODO: Implement Authorization middleware, then create GetMidwife, UpdateMidwife, and DeleteMidwife functions.
