package models

import "time"

type Review struct {
	ID             uint      `gorm:"primaryKey"`
	CreatedAt      time.Time `gorm:"not null"`
	TrainingPlanID uint      `gorm:"not null;uniqueIndex:idx_user_id_training_plan_id"`
	UserID         string    `gorm:"not null;uniqueIndex:idx_user_id_training_plan_id"`
	Score          uint      `gorm:"not null"`
	Comment        string
}
