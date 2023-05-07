package models

import "time"

type TrainingPlan struct {
	ID          uint   `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"not null"`
	Duration    uint   `gorm:"not null"`
	TrainerID   string
	CreatedAt   time.Time  `gorm:"not null"`
	Exercises   []Exercise `gorm:"foreignKey:TrainingPlanID"`
}
