package goals

import (
	"time"

	"github.com/fiufit/trainings/contracts"
)

type UpdateGoalRequest struct {
	Title     string    `json:"title" binding:"required"`
	GoalValue uint      `json:"value" binding:"required"`
	Deadline  time.Time `json:"deadline" binding:"required"`
	UserID    string    `json:"user_id" binding:"required"`
	GoalID    uint
}

func (req *UpdateGoalRequest) Validate() error {
	if req.Deadline.Before(time.Now()) {
		return contracts.ErrBadRequest
	}
	return nil
}
