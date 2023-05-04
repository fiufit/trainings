package usecases

import (
	"context"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingGetter interface {
	GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error)
}

type TrainingGetterImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewTrainingGetterImpl(trainings repositories.TrainingPlans, logger *zap.Logger) TrainingGetterImpl {
	return TrainingGetterImpl{trainings: trainings, logger: logger}
}

func (uc *TrainingGetterImpl) GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error) {
	return uc.trainings.GetTrainingPlans(ctx, req)
}
