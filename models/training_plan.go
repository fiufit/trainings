package models

import (
	"time"

	"gorm.io/gorm"
)

type TrainingPlan struct {
	ID          uint   `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"not null"`
	Duration    uint   `gorm:"not null"`
	TrainerID   string
	CreatedAt   time.Time `gorm:"not null"`
	DeletedAt   gorm.DeletedAt
	Exercises   []Exercise `gorm:"foreignKey:TrainingPlanID"`
	Reviews     []Review   `gorm:"foreignKey:TrainingPlanID" json:"-"`
	MeanScore   float32    `gorm:"-"`
	PictureUrl  string     `gorm:"-"`
}
