package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingUpdater interface {
	UpdateTrainingPlan(ctx context.Context, req trainings.UpdateTrainingRequest) (models.TrainingPlan, error)
}

type TrainingUpdaterImpl struct {
	trainings repositories.TrainingPlans
	firebase  repositories.Firebase
	logger    *zap.Logger
}

func NewTrainingUpdaterImpl(trainings repositories.TrainingPlans, firebase repositories.Firebase, logger *zap.Logger) TrainingUpdaterImpl {
	return TrainingUpdaterImpl{trainings: trainings, firebase: firebase, logger: logger}
}

func (uc *TrainingUpdaterImpl) UpdateTrainingPlan(ctx context.Context, req trainings.UpdateTrainingRequest) (models.TrainingPlan, error) {
	oldTraining, err := uc.getTrainingPlan(ctx, req.ID, req.TrainerID)
	if err != nil {
		return models.TrainingPlan{}, err
	}

	training := trainings.ConverToTrainingPlan(req.BaseTrainingRequest)
	training.ID = oldTraining.ID

	uc.firebase.FillTrainingPicture(ctx, &training)

	updatedTraining, err := uc.trainings.UpdateTrainingPlan(ctx, training)
	if err != nil {
		return models.TrainingPlan{}, err
	}
	return updatedTraining, nil
}

func (uc *TrainingUpdaterImpl) getTrainingPlan(ctx context.Context, trainingID uint, trainerID string) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return models.TrainingPlan{}, err
	}
	if training.TrainerID != trainerID {
		return models.TrainingPlan{}, contracts.ErrUnauthorizedTrainer
	}
	return training, nil
}
