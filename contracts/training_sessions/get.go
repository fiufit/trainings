package training_sessions

import (
	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
)

type GetTrainingSessionsRequest struct {
	UserID     string `form:"user_id" binding:"required"`
	TrainingID uint   `form:"training_id"`
	contracts.Pagination
}

type GetTrainingSessionsResponse struct {
	Sessions []models.TrainingSession `json:"training_sessions"`
	contracts.Pagination
}
