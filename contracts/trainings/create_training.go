package trainings

import (
	"time"

	"github.com/fiufit/trainings/models"
)

type CreateTrainingRequest struct {
	BaseTrainingRequest
}

type BaseTrainingRequest struct {
	TrainerID   string            `json:"trainer_id" binding:"required"`
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	Difficulty  string            `json:"difficulty" binding:"required"`
	Duration    uint              `json:"duration" binding:"required"`
	Exercises   []ExerciseRequest `json:"exercises" binding:"required"`
	Tags        []models.Tag      `json:"-"`
	TagStrings  []string          `json:"tags"`
}

func (req *CreateTrainingRequest) Validate() error {
	tags, err := models.ValidateTags(req.TagStrings...)
	if err != nil {
		return err
	}
	req.Tags = tags
	return nil
}

type CreateTrainingResponse struct {
	TrainingPlan models.TrainingPlan `json:"training_plan"`
}

type ExerciseRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func ConvertToExercise(exerciseReq ExerciseRequest) models.Exercise {
	return models.Exercise{
		Title:          exerciseReq.Title,
		Description:    exerciseReq.Description,
		ID:             0,
		TrainingPlanID: 0,
	}
}

func ConvertToExercises(exerciseReqs []ExerciseRequest) []models.Exercise {
	exercises := make([]models.Exercise, len(exerciseReqs))
	for i, exerciseReq := range exerciseReqs {
		exercises[i] = ConvertToExercise(exerciseReq)
	}
	return exercises
}

func ConverToTrainingPlan(trainingReq BaseTrainingRequest) models.TrainingPlan {
	exercises := ConvertToExercises(trainingReq.Exercises)
	return models.TrainingPlan{
		Name:        trainingReq.Name,
		Description: trainingReq.Description,
		Difficulty:  trainingReq.Difficulty,
		TrainerID:   trainingReq.TrainerID,
		Duration:    trainingReq.Duration,
		CreatedAt:   time.Now(),
		Exercises:   exercises,
		Tags:        trainingReq.Tags,
	}
}
