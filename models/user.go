package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                string
	Nickname          string
	DisplayName       string
	IsMale            bool
	CreatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	BornAt            time.Time
	Height            uint
	Weight            uint
	IsVerifiedTrainer bool
	MainLocation      string
	Interests         []string
	PictureUrl        string
}
