package reviews

import "github.com/fiufit/trainings/contracts"

type UpdateReviewRequest struct {
	Score          uint   `json:"score"`
	Comment        string `json:"comment"`
	UserID         string `json:"user_id" binding:"required"`
	TrainingPlanID uint
	ReviewID       uint
}

func (req *UpdateReviewRequest) Validate() error {
	if req.Score < 1 || req.Score > 5 {
		return contracts.ErrBadRequest
	}
	return nil
}
