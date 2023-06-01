package training_sessions

import "github.com/fiufit/trainings/models"

type CreateTrainingSessionResponse struct {
	Session models.TrainingSession `json:"training_session"`
}
