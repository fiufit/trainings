package training

import "github.com/fiufit/trainings/models"

type CreateTrainingRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	TrainerID   string            `json:"trainer_id" binding:"required"`
	Exercises   []ExerciseRequest `json:"exercises" binding:"required"`
}

type CreateTrainingResponse struct {
	TrainingPlan models.TrainingPlan `json:"training_plan"`
}

type ExerciseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
