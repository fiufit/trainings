package goals

import (
	"time"

	"github.com/fiufit/trainings/contracts"
)

type UpdateGoalRequest struct {
	Title     string    `json:"title"`
	GoalValue uint      `json:"value"`
	Deadline  time.Time `json:"deadline"`
	UserID    string    `json:"user_id" binding:"required"`
	GoalID    uint
}

func (req *UpdateGoalRequest) Validate() error {
	if !req.Deadline.IsZero() && req.Deadline.Before(time.Now()) {
		return contracts.ErrBadRequest
	}
	return nil
}
