package training

import (
	"time"

	"github.com/fiufit/trainings/models"
)

type CreateTrainingRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description" binding:"required"`
	Difficulty  string            `json:"difficulty" binding:"required"`
	TrainerID   string            `json:"trainer_id" binding:"required"`
	Duration    int8              `json:"duration" binding:"required"`
	Exercises   []ExerciseRequest `json:"exercises" binding:"required"`
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
		Done:           false,
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

func ConverToTrainingPlan(trainingReq CreateTrainingRequest) models.TrainingPlan {
	exercises := ConvertToExercises(trainingReq.Exercises)
	return models.TrainingPlan{
		Name:        trainingReq.Name,
		Description: trainingReq.Description,
		Difficulty:  trainingReq.Difficulty,
		TrainerID:   trainingReq.TrainerID,
		Duration:    trainingReq.Duration,
		CreatedAt:   time.Now(),
		Exercises:   exercises,
	}
}
