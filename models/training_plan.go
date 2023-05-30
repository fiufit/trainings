package models

import (
	"time"

	"gorm.io/gorm"
)

type TrainingPlan struct {
	ID          uint   `gorm:"primaryKey;not null"`
	Version     uint   `gorm:"primaryKey;not null;autoincrement:false"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"not null"`
	Duration    uint   `gorm:"not null"`
	TrainerID   string
	CreatedAt   time.Time `gorm:"not null"`
	DeletedAt   gorm.DeletedAt
	Exercises   []Exercise
	Tags        []Tag `gorm:"many2many:training_plan_tags"`
	Reviews     []Review
	MeanScore   float32 `gorm:"-"`
	PictureUrl  string  `gorm:"-"`
}
