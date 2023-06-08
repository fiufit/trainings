package models

import "time"

type Goal struct {
	ID                uint      `gorm:"primaryKey"`
	Title             string    `json:"title" gorm:"not null"`
	GoalValue         uint      `json:"value" gorm:"not null"`
	GoalValueProgress uint      `json:"progress" gorm:"not null"`
	GoalType          string    `json:"type" gorm:"not null"`
	GoalSubtype       string    `json:"subtype"`
	CreatedAt         time.Time `json:"created_at" gorm:"not null"`
	Deadline          time.Time `json:"deadline" gorm:"not null"`
	UserID            string    `json:"user_id" gorm:"not null"`
}
