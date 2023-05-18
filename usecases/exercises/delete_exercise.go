package exercises

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ExerciseDeleter interface {
	DeleteExercise(ctx context.Context, req exercises.DeleteExerciseRequest) error
}

type ExerciseDeleterImpl struct {
	trainings repositories.TrainingPlans
	exercises repositories.Exercises
	logger    *zap.Logger
}

func NewExerciseDeleterImpl(trainings repositories.TrainingPlans, exercises repositories.Exercises, logger *zap.Logger) ExerciseDeleterImpl {
	return ExerciseDeleterImpl{trainings: trainings, exercises: exercises, logger: logger}
}

func (uc *ExerciseDeleterImpl) DeleteExercise(ctx context.Context, req exercises.DeleteExerciseRequest) error {
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return err
	}
	if training.TrainerID != req.TrainerID {
		return contracts.ErrUnauthorizedTrainer
	}
	return uc.exercises.DeleteExercise(ctx, req.ExerciseID)
}
