package trainings

import "github.com/fiufit/trainings/models"

type UpdateTrainingRequest struct {
	ID          uint
	TrainerID   string `json:"trainer_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Duration    uint   `json:"duration"`
}

type UpdateTrainingResponse struct {
	TrainingPlan models.TrainingPlan `json:"training_plan"`
}
