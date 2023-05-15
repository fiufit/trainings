package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingDeleter interface {
	DeleteTraining(ctx context.Context, req trainings.DeleteTrainingRequest) error
}

type TrainingDeleterImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewTrainingDeleterImpl(trainings repositories.TrainingPlans, logger *zap.Logger) TrainingDeleterImpl {
	return TrainingDeleterImpl{trainings: trainings, logger: logger}
}

func (uc *TrainingDeleterImpl) DeleteTraining(ctx context.Context, req trainings.DeleteTrainingRequest) error {
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return err
	}
	if training.TrainerID != req.TrainerID {
		return contracts.ErrUnauthorizedTrainer
	}
	return uc.trainings.DeleteTrainingPlan(ctx, req.TrainingPlanID)
}
