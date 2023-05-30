package trainings

import (
	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
)

type GetTrainingsRequest struct {
	Name        string   `form:"name"`
	Description string   `form:"description"`
	Difficulty  string   `form:"difficulty"`
	TrainerID   string   `form:"trainer_id"`
	MinDuration uint     `form:"min_duration"`
	MaxDuration uint     `form:"max_duration"`
	UserID      string   `form:"user_id"`
	TagStrings  []string `form:"tags[]"`
	Tags        []models.Tag
	contracts.Pagination
}

func (req *GetTrainingsRequest) Validate() error {
	tags, err := models.ValidateTags(req.TagStrings...)
	if err != nil {
		return err
	}
	req.Tags = tags
	return nil
}

type GetTrainingsResponse struct {
	Pagination    contracts.Pagination  `json:"pagination"`
	TrainingPlans []models.TrainingPlan `json:"trainings"`
}
