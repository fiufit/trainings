package models

import "time"

type TrainingPlan struct {
	ID          int8   `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Difficulty  string `gorm:"not null"`
	Duration    int8   `gorm:"not null"`
	TrainerID   string
	Trainer     User       `gorm:"foreignKey:TrainerID" json:"-"`
	CreatedAt   time.Time  `gorm:"not null"`
	Exercises   []Exercise `gorm:"foreignKey:TrainingPlanID"`
}
