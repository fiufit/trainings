package models

type Exercise struct {
	ID             uint   `gorm:"primaryKey"`
	TrainingPlanID uint   `gorm:"not null"`
	Title          string `gorm:"not null"`
	Description    string `gorm:"not null"`
}
