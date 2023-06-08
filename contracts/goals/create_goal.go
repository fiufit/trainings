package goals

import (
	"time"

	"github.com/fiufit/trainings/contracts"
)

type CreateGoalRequest struct {
	Title       string    `json:"title" binding:"required"`
	GoalType    string    `json:"type" binding:"required"`
	GoalSubtype string    `json:"subtype"`
	GoalValue   uint      `json:"value" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	UserID      string    `json:"user_id" binding:"required"`
}

func (req *CreateGoalRequest) Validate() error {
	if req.Deadline.Before(time.Now()) {
		return contracts.ErrBadRequest
	}
	return nil
}
