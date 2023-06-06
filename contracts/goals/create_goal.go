package goals

import "time"

type CreateGoalRequest struct {
	Title       string    `json:"title" binding:"required"`
	GoalType    string    `json:"type" binding:"required"`
	GoalSubtype string    `json:"subtype"`
	GoalValue   uint      `json:"value" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	UserID      string    `json:"user_id" binding:"required"`
}
