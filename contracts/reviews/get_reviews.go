package reviews

import (
	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
)

type GetReviewsRequest struct {
	MinScore       uint   `form:"min_score"`
	MaxScore       uint   `form:"max_score"`
	Comment        string `form:"comment"`
	TrainingPlanID uint
	UserID         string `form:"user_id"`
	contracts.Pagination
}

type GetReviewsResponse struct {
	Pagination contracts.Pagination `json:"pagination"`
	Reviews    []models.Review      `json:"reviews"`
}
