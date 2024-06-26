package database

import "github.com/muhreeowki/midwifery-ehr/internal/models"

/* CreateMidwife creates a new midwife record in the database. */
func (engine *DatabaseEngine) CreateMidwife(midwifeData *models.Midwife) (midwife models.Midwife, err error) {
	midwife = *midwifeData
	err = engine.DB.Create(midwife).Error
	if err != nil {
		return midwife, err
	}
	return midwife, nil
}
