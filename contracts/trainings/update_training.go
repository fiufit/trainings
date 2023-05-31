package trainings

import "github.com/fiufit/trainings/models"

type UpdateTrainingRequest struct {
	ID uint
	BaseTrainingRequest
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
