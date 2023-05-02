package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                string    `gorm:"primaryKey;not null"`
	Nickname          string    `gorm:"not null;unique;index"`
	DisplayName       string    `gorm:"not null"`
	IsMale            bool      `gorm:"not null"`
	CreatedAt         time.Time `gorm:"not null"`
	DeletedAt         gorm.DeletedAt
	BornAt            time.Time `gorm:"not null"`
	Height            uint      `gorm:"not null"`
	Weight            uint      `gorm:"not null"`
	IsVerifiedTrainer bool      `gorm:"not null;default:false"`
	MainLocation      string    `gorm:"not null"`
	Interests         []string  `gorm:"-"`
}
