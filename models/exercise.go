package models

type Exercise struct {
	ID             int8   `gorm:"primaryKey"`
	TrainingPlanID int8   `gorm:"not null"`
	Title          string `gorm:"not null"`
	Description    string `gorm:"not null"`
}
