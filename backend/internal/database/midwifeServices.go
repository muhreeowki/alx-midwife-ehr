package database

import "github.com/muhreeowki/midwifery-ehr/internal/models"

/* CreateMidwife creates a new midwife record in the database. */
func (engine *DatabaseEngine) CreateMidwife(midwife *models.Midwife) (err error) {
	err = engine.DB.Create(midwife).Error
	return err
}
