package trainings

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
	users     repositories.Users
	logger    *zap.Logger
}

func NewTrainingCreatorImpl(trainings repositories.TrainingPlans, users repositories.Users, logger *zap.Logger) TrainingCreatorImpl {
	return TrainingCreatorImpl{trainings: trainings, users: users, logger: logger}
}

func (uc *TrainingCreatorImpl) CreateTraining(ctx context.Context, req training.CreateTrainingRequest) (training.CreateTrainingResponse, error) {
	_, err := uc.users.GetUserByID(ctx, req.TrainerID)
	if err != nil {
		return training.CreateTrainingResponse{}, err
	}
	newTraining := training.ConverToTrainingPlan(req)
	createdTraining, err := uc.trainings.CreateTrainingPlan(ctx, newTraining)
	return training.CreateTrainingResponse{TrainingPlan: createdTraining}, err
}
