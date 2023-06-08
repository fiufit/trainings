package goals

import (
	"time"

	"github.com/fiufit/trainings/models"
)

type GetGoalsRequest struct {
	UserID      string    `form:"user_id" binding:"required"`
	GoalType    string    `form:"type"`
	GoalSubtype string    `form:"subtype"`
	Deadline    time.Time `form:"deadline"`
}

type GetGoalsResponse struct {
	Goals []models.Goal `json:"goals"`
}
