package helpers

import (
	"time"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
)

func ConvertToExercise(exerciseReq training.ExerciseRequest) models.Exercise {
	return models.Exercise{
		Title:          exerciseReq.Title,
		Description:    exerciseReq.Description,
		Done:           false,
		ID:             0,
		TrainingPlanID: 0,
	}
}

func ConvertToExercises(exerciseReqs []training.ExerciseRequest) []models.Exercise {
	exercises := make([]models.Exercise, len(exerciseReqs))
	for i, exerciseReq := range exerciseReqs {
		exercises[i] = ConvertToExercise(exerciseReq)
	}
	return exercises
}

func ConverToTrainingPlan(trainingReq training.CreateTrainingRequest) models.TrainingPlan {
	exercises := ConvertToExercises(trainingReq.Exercises)
	return models.TrainingPlan{
		Name:        trainingReq.Name,
		Description: trainingReq.Description,
		TrainerID:   trainingReq.TrainerID,
		CreatedAt:   time.Now(),
		Exercises:   exercises,
	}
}
