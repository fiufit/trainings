package models

type Favorite struct {
	ID                  uint `gorm:"primaryKey"`
	TrainingPlanID      uint `gorm:"not null;uniqueIndex:fav_idx_user_id_training_plan_id_training_plan_version"`
	TrainingPlan        TrainingPlan
	TrainingPlanVersion uint   `gorm:"not null"`
	UserID              string `gorm:"not null;uniqueIndex:fav_idx_user_id_training_plan_id_training_plan_version" json:"-"`
}
