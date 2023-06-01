package models

import "time"

type ExerciseSession struct {
	ID                uint `gorm:"primaryKey;not null"`
	TrainingSessionID uint `gorm:"not null"`
	Done              bool `gorm:"not null"`
	ExerciseID        uint `gorm:"not null"`
	Exercise          Exercise
}

type TrainingSession struct {
	ID                  uint         `gorm:"primaryKey;not null"`
	TrainingPlanID      uint         `gorm:"not null"`
	TrainingPlanVersion uint         `gorm:"not null"`
	TrainingPlan        TrainingPlan `gorm:"ForeignKey:TrainingPlanID,TrainingPlanVersion;References:ID,Version"`
	UserID              string       `gorm:"not null"`
	Done                bool         `gorm:"not null"`
	StepCount           uint         `gorm:"not null"`
	SecondsCount        uint         `gorm:"not null"`
	ExerciseSessions    []ExerciseSession
	UpdatedAt           time.Time
}
