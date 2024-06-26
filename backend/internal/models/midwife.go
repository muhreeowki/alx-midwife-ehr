package models

import (
	"gorm.io/gorm"
)

type Midwife struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null;" binding:"required"`
	Email    string `json:"email" gorm:"not null;unique" validate:"email" binding:"required"`
	Password string `json:"password" gorm:"not null;" binding:"required"`
	Patients []Patient
}
