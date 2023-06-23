package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts/metrics"
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
	metrics   repositories.Metrics
	logger    *zap.Logger
}

func NewTrainingCreatorImpl(trainings repositories.TrainingPlans, users repositories.Users, metrics repositories.Metrics, logger *zap.Logger) TrainingCreatorImpl {
	return TrainingCreatorImpl{trainings: trainings, users: users, metrics: metrics, logger: logger}
}

func (uc *TrainingCreatorImpl) CreateTraining(ctx context.Context, req trainings.CreateTrainingRequest) (trainings.CreateTrainingResponse, error) {
	trainer, err := uc.users.GetUserByID(ctx, req.TrainerID)
	if err != nil {
		return trainings.CreateTrainingResponse{}, err
	}
	newTraining := trainings.ConverToTrainingPlan(req.BaseTrainingRequest)
	createdTraining, err := uc.trainings.CreateTrainingPlan(ctx, newTraining)

	if err != nil {
		return trainings.CreateTrainingResponse{}, err
	}

	createdTrainingMetric := metrics.CreateMetricRequest{
		MetricType: "new_training",
		SubType:    trainer.ID,
	}
	uc.metrics.Create(ctx, createdTrainingMetric)

	for _, tag := range createdTraining.Tags {
		taggedTrainingMetric := metrics.CreateMetricRequest{
			MetricType: "training_tagged",
			SubType:    tag.Name,
		}
		uc.metrics.Create(ctx, taggedTrainingMetric)
	}

	return trainings.CreateTrainingResponse{TrainingPlan: createdTraining}, nil
}
