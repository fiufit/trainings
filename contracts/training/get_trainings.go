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
	Duration    int8   `form:"duration"`
	contracts.Pagination
}

type GetTrainingsResponse struct {
	Pagination    contracts.Pagination  `json:"pagination"`
	TrainingPlans []models.TrainingPlan `json:"trainings"`
}
