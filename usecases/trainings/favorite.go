package trainings

import (
	"context"

	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type FavoriteAdder interface {
	AddToFavorite(ctx context.Context, userID string, trainingID uint) error
	RemoveFromFavorite(ctx context.Context, userID string, trainingID uint) error
}

type FavoriteAdderImpl struct {
	trainings repositories.TrainingPlans
	logger    *zap.Logger
}

func NewFavoriteAdderImpl(trainings repositories.TrainingPlans, logger *zap.Logger) FavoriteAdderImpl {
	return FavoriteAdderImpl{trainings: trainings, logger: logger}
}

func (uc *FavoriteAdderImpl) AddToFavorite(ctx context.Context, userID string, trainingID uint) error {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return err
	}
	return uc.trainings.AddToFavorite(ctx, userID, trainingID, training.Version)
}

func (uc *FavoriteAdderImpl) RemoveFromFavorite(ctx context.Context, userID string, trainingID uint) error {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return err
	}
	return uc.trainings.RemoveFromFavorite(ctx, userID, trainingID, training.Version)
}
