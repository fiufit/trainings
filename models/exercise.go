package models

import (
	"gorm.io/gorm"
)

type Exercise struct {
	ID                  uint `gorm:"primaryKey"`
	TrainingPlanID      uint `gorm:"not null"`
	TrainingPlanVersion uint `gorm:"not null"`
	DeletedAt           gorm.DeletedAt
	Title               string `gorm:"not null"`
	Description         string `gorm:"not null"`
}
