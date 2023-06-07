package goals

import "github.com/fiufit/trainings/models"

type GetGoalsRequest struct {
	UserID string `form:"user_id"`
}

type GetGoalsResponse struct {
	Goals []models.Goal `json:"goals"`
}
