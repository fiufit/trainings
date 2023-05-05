package training

import "github.com/fiufit/trainings/models"

type UpdateTrainingRequest struct {
	ID          string
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Duration    uint   `json:"duration"`
}

type UpdateTrainingResponse struct {
	TrainingPlan models.TrainingPlan `json:"training_plan"`
}
