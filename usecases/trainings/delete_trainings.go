package trainings

import (
	"context"
	"strconv"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingDeleter interface {
	DeleteTraining(ctx context.Context, req training.DeleteTrainingRequest) error
}

type TrainingDeleterImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewTrainingDeleterImpl(trainings repositories.TrainingPlans, logger *zap.Logger) TrainingDeleterImpl {
	return TrainingDeleterImpl{trainings: trainings, logger: logger}
}

func (uc *TrainingDeleterImpl) DeleteTraining(ctx context.Context, req training.DeleteTrainingRequest) error {
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return err
	}
	if training.TrainerID != req.TrainerID {
		return contracts.ErrUnauthorizedTrainer
	}
	u64, err := strconv.ParseUint(req.TrainingPlanID, 10, 32)
	if err != nil {
		return err
	}
	trainingID := uint(u64)
	return uc.trainings.DeleteTrainingPlan(ctx, trainingID)
}
