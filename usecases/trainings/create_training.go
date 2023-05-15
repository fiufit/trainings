package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingCreator interface {
	CreateTraining(ctx context.Context, req trainings.CreateTrainingRequest) (trainings.CreateTrainingResponse, error)
}

type TrainingCreatorImpl struct {
	trainings repositories.TrainingPlans
	users     repositories.Users
	logger    *zap.Logger
}

func NewTrainingCreatorImpl(trainings repositories.TrainingPlans, users repositories.Users, logger *zap.Logger) TrainingCreatorImpl {
	return TrainingCreatorImpl{trainings: trainings, users: users, logger: logger}
}

func (uc *TrainingCreatorImpl) CreateTraining(ctx context.Context, req trainings.CreateTrainingRequest) (trainings.CreateTrainingResponse, error) {
	_, err := uc.users.GetUserByID(ctx, req.TrainerID)
	if err != nil {
		return trainings.CreateTrainingResponse{}, err
	}
	newTraining := trainings.ConverToTrainingPlan(req)
	createdTraining, err := uc.trainings.CreateTrainingPlan(ctx, newTraining)
	return trainings.CreateTrainingResponse{TrainingPlan: createdTraining}, err
}
