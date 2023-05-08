package users

import (
	"time"

	"gorm.io/gorm"
)

type GetUserResponse struct {
	ID                string         `json:"ID"`
	Nickname          string         `json:"Nickname"`
	DisplayName       string         `json:"DisplayName"`
	IsMale            bool           `json:"IsMale"`
	CreatedAt         time.Time      `json:"CreatedAt"`
	DeletedAt         gorm.DeletedAt `json:"DeletedAt"`
	BornAt            time.Time      `json:"BornAt"`
	Height            uint           `json:"Height"`
	Weight            uint           `json:"Weight"`
	IsVerifiedTrainer bool           `json:"IsVerifiedTrainer"`
	MainLocation      string         `json:"MainLocation"`
	Interests         []string       `json:"Interests"`
	PictureUrl        string         `json:"PictureUrl"`
}
