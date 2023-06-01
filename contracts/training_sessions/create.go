package training_sessions

import (
	"github.com/fiufit/trainings/models"
)

type CreateTrainingSessionRequest struct {
	UserID     string `form:"user_id" binding:"required"`
	TrainingID uint   `form:"training_id" binding:"required"`
}

type CreateTrainingSessionResponse struct {
	Session models.TrainingSession `json:"training_session"`
}
