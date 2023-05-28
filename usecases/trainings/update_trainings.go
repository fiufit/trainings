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
	training, err := uc.getTrainingPlan(ctx, req.ID, req.TrainerID)
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

func (uc *TrainingUpdaterImpl) patchTrainingModel(ctx context.Context, training models.TrainingPlan, req trainings.UpdateTrainingRequest) (models.TrainingPlan, error) {
	training.Tags = req.Tags

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

func (uc *TrainingUpdaterImpl) getTrainingPlan(ctx context.Context, trainingID uint, trainerID string) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return models.TrainingPlan{}, err
	}
	if training.TrainerID != trainerID {
		return models.TrainingPlan{}, contracts.ErrUnauthorizedTrainer
	}
	trainingPictureUrl := uc.firebase.GetTrainingPictureUrl(ctx, trainingID, trainerID)
	training.PictureUrl = trainingPictureUrl

	sum := uint(0)
	for i := range training.Reviews {
		sum += training.Reviews[i].Score
	}
	reviews := uint(len(training.Reviews))
	if reviews != 0 {
		training.MeanScore = float32(sum / reviews)
	} else {
		training.MeanScore = 0
	}
	return training, nil
}
