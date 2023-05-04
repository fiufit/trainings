package training

import (
	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
)

type GetTrainingsRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Difficulty  string `form:"difficulty"`
	TrainerID   string `form:"trainer_id"`
	MinDuration int8   `form:"min_duration"`
	MaxDuration int8   `form:"max_duration"`
	contracts.Pagination
}

type GetTrainingsResponse struct {
	Pagination    contracts.Pagination  `json:"pagination"`
	TrainingPlans []models.TrainingPlan `json:"trainings"`
}
