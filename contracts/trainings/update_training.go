package trainings

import "github.com/fiufit/trainings/models"

type UpdateTrainingRequest struct {
	ID          uint
	TrainerID   string       `json:"trainer_id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Difficulty  string       `json:"difficulty"`
	Duration    uint         `json:"duration"`
	TagStrings  []string     `json:"tags"`
	Tags        []models.Tag `json:"-"`
}

type UpdateTrainingResponse struct {
	TrainingPlan models.TrainingPlan `json:"training_plan"`
}

func (req *UpdateTrainingRequest) Validate() error {
	tags, err := models.ValidateTags(req.TagStrings...)
	if err != nil {
		return err
	}
	req.Tags = tags
	return nil
}
