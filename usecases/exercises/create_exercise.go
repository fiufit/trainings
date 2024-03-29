package exercises

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ExerciseCreator interface {
	CreateExercise(ctx context.Context, req exercises.CreateExerciseRequest) (models.Exercise, error)
}

type ExerciseCreatorImpl struct {
	trainings repositories.TrainingPlans
	exercises repositories.Exercises
	logger    *zap.Logger
}

func NewExerciseCreatorImpl(trainings repositories.TrainingPlans, exercises repositories.Exercises, logger *zap.Logger) ExerciseCreatorImpl {
	return ExerciseCreatorImpl{trainings: trainings, exercises: exercises, logger: logger}
}

func (uc *ExerciseCreatorImpl) CreateExercise(ctx context.Context, req exercises.CreateExerciseRequest) (models.Exercise, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return models.Exercise{}, err
	}
	if training.TrainerID != req.TrainerID {
		return models.Exercise{}, contracts.ErrUnauthorizedTrainer
	}
	newExercise := models.Exercise{
		TrainingPlanID: req.TrainingPlanID,
		Title:          req.Title,
		Description:    req.Description,
	}
	createdExercise, err := uc.exercises.CreateExercise(ctx, newExercise)
	return createdExercise, err
}
