package reviews

import (
	"github.com/fiufit/trainings/contracts"
)

type CreateReviewRequest struct {
	Score          uint   `json:"score" binding:"required"`
	Comment        string `json:"comment"`
	UserID         string `json:"user_id" binding:"required"`
	TrainingPlanID uint
}

func (req *CreateReviewRequest) Validate() error {
	if req.Score < 1 || req.Score > 5 {
		return contracts.ErrBadRequest
	}
	return nil
}
