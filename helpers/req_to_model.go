package helpers

import (
	"time"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
)

func convertToExercise(exerciseReq training.ExerciseRequest) models.Exercise {
	return models.Exercise{
		Title:          exerciseReq.Title,
		Description:    exerciseReq.Description,
		Done:           false,
		ID:             0,
		TrainingPlanID: 0,
	}
}

func convertToExercises(exerciseReqs []training.ExerciseRequest) []models.Exercise {
	exercises := make([]models.Exercise, len(exerciseReqs))
	for i, exerciseReq := range exerciseReqs {
		exercises[i] = convertToExercise(exerciseReq)
	}
	return exercises
}

func converToTrainingPlan(trainingReq training.CreateTrainingRequest, exercises []models.Exercise) models.TrainingPlan {
	return models.TrainingPlan{
		Name:        trainingReq.Name,
		Description: trainingReq.Description,
		TrainerID:   trainingReq.TrainerID,
		CreatedAt:   time.Now(),
		Exercises:   exercises,
	}
}
