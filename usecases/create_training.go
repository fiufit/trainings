package usecases

import (
	"context"

	"github.com/fiufit/trainings/contracts/training"
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
	newTraining := training.ConverToTrainingPlan(req)
	createdTraining, err := uc.trainings.CreateTrainingPlan(ctx, newTraining)
	return training.CreateTrainingResponse{TrainingPlan: createdTraining}, err
}
