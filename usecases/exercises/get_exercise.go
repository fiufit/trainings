package exercises

import (
	"context"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ExerciseGetter interface {
	GetExerciseByID(ctx context.Context, req training.GetExerciseRequest) (models.Exercise, error)
}

type ExerciseGetterImpl struct {
	trainings repositories.TrainingPlans
	exercises repositories.Exercises
	logger    *zap.Logger
}

func NewExerciseGetterImpl(trainings repositories.TrainingPlans, exercises repositories.Exercises, logger *zap.Logger) ExerciseGetterImpl {
	return ExerciseGetterImpl{trainings: trainings, exercises: exercises, logger: logger}
}

func (uc *ExerciseGetterImpl) GetExerciseByID(ctx context.Context, req training.GetExerciseRequest) (models.Exercise, error) {
	_, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return models.Exercise{}, err
	}
	return uc.exercises.GetExerciseByID(ctx, req.ExerciseID)
}
