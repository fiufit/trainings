package usecases

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ExerciseUpdater interface {
	UpdateExercise(ctx context.Context, req training.UpdateExerciseRequest) (models.Exercise, error)
}

type ExerciseUpdaterImpl struct {
	trainings repositories.TrainingPlans
	exercises repositories.Exercises
	logger    *zap.Logger
}

func NewExerciseUpdaterImpl(trainings repositories.TrainingPlans, exercises repositories.Exercises, logger *zap.Logger) ExerciseUpdaterImpl {
	return ExerciseUpdaterImpl{trainings: trainings, exercises: exercises, logger: logger}
}

func (uc *ExerciseUpdaterImpl) UpdateExercise(ctx context.Context, req training.UpdateExerciseRequest) (models.Exercise, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return models.Exercise{}, err
	}
	if training.TrainerID != req.TrainerID {
		return models.Exercise{}, contracts.ErrUnauthorizedTrainer
	}

	exercise, err := uc.exercises.GetExerciseByID(ctx, req.ExerciseID)

	patchedExercise, err := uc.patchExerciseModel(ctx, exercise, req)
	if err != nil {
		return models.Exercise{}, err
	}

	updatedExercise, err := uc.exercises.UpdateExercise(ctx, patchedExercise)
	if err != nil {
		return models.Exercise{}, err
	}
	return updatedExercise, nil
}

func (uc *ExerciseUpdaterImpl) patchExerciseModel(ctx context.Context, exercise models.Exercise, req training.UpdateExerciseRequest) (models.Exercise, error) {
	if req.Title != "" {
		exercise.Title = req.Title
	}

	if req.Description != "" {
		exercise.Description = req.Description
	}

	return exercise, nil
}
