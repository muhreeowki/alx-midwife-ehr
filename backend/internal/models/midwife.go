package models

import (
	"gorm.io/gorm"
)

type Midwife struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"not null;" binding:"required"`
	LastName  string `json:"last_name" gorm:"not null;" binding:"required"`
	Email     string `json:"email" gorm:"not null;unique" validate:"email" binding:"required"`
	Password  string `json:"password" gorm:"not null;" binding:"required"`
	ImageURL  string `json:"profile_image"`
	Patients  []Patient
}
