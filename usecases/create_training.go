package usecases

import (
	"context"
	"time"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingCreator interface {
	CreateTraining(ctx context.Context, req training.CreateTrainingRequest) (training.CreateTrainingResponse, error)
}

type TrainingCreatorImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewTrainingCreatorImpl(trainings repositories.TrainingPlans, logger *zap.Logger) TrainingCreatorImpl {
	return TrainingCreatorImpl{trainings: trainings, logger: logger}
}

func (uc *TrainingCreatorImpl) CreateTraining(ctx context.Context, req training.CreateTrainingRequest) (training.CreateTrainingResponse, error) {
	exercises := convertToExercises(req.Exercises)
	newTraining := models.TrainingPlan{
		Name:        req.Name,
		Description: req.Description,
		TrainerID:   req.TrainerID,
		CreatedAt:   time.Now(),
		Exercises:   exercises,
	}
	createdTraining, err := uc.trainings.CreateTrainingPlan(ctx, newTraining)
	return training.CreateTrainingResponse{TrainingPlan: createdTraining}, err
}

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
