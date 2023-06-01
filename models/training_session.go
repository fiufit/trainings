package models

import "time"

type ExerciseSession struct {
	ID                uint
	TrainingSessionID uint
	ExerciseID        uint
	Exercise          Exercise
	Done              bool
}

type TrainingSession struct {
	ID                  uint         `gorm:"primaryKey;not null"`
	TrainingPlanID      uint         `gorm:"not null"`
	TrainingPlanVersion uint         `gorm:"not null"`
	TrainingPlan        TrainingPlan `gorm:"ForeignKey:TrainingPlanID,TrainingPlanVersion;References:ID,Version"`
	UserID              string
	ExerciseSessions    []ExerciseSession
	Done                bool
	StepCount           uint
	SecondsCount        uint
	UpdatedAt           time.Time
}
