package models

import "time"

type Review struct {
	ID                  uint      `gorm:"primaryKey"`
	CreatedAt           time.Time `gorm:"not null"`
	TrainingPlanID      uint      `gorm:"not null;uniqueIndex:idx_user_id_training_plan_id_training_plan_version"`
	TrainingPlanVersion uint      `gorm:"not null;uniqueIndex:idx_user_id_training_plan_id_training_plan_version"`
	UserID              string    `gorm:"not null;uniqueIndex:idx_user_id_training_plan_id_training_plan_version" json:"-"`
	User                User      `gorm:"-"`
	Score               uint      `gorm:"not null"`
	Comment             string
}
