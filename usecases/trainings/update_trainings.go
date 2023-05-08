package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingUpdater interface {
	UpdateTrainingPlan(ctx context.Context, req training.UpdateTrainingRequest) (models.TrainingPlan, error)
}

type TrainingUpdaterImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewTrainingUpdaterImpl(trainings repositories.TrainingPlans, logger *zap.Logger) TrainingUpdaterImpl {
	return TrainingUpdaterImpl{trainings: trainings, logger: logger}
}

func (uc *TrainingUpdaterImpl) UpdateTrainingPlan(ctx context.Context, req training.UpdateTrainingRequest) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, req.ID)
	if err != nil {
		return models.TrainingPlan{}, err
	}

	patchedTraining, err := uc.patchTrainingModel(ctx, training, req)
	if err != nil {
		return models.TrainingPlan{}, err
	}

	updatedTraining, err := uc.trainings.UpdateTrainingPlan(ctx, patchedTraining)
	if err != nil {
		return models.TrainingPlan{}, err
	}
	return updatedTraining, nil
}

func (uc *TrainingUpdaterImpl) patchTrainingModel(ctx context.Context, training models.TrainingPlan, req training.UpdateTrainingRequest) (models.TrainingPlan, error) {
	if req.Name != "" {
		training.Name = req.Name
	}

	if req.Description != "" {
		training.Description = req.Description
	}

	if req.Difficulty != "" {
		training.Difficulty = req.Difficulty
	}

	if req.Duration != 0 {
		training.Duration = req.Duration
	}

	return training, nil
}
