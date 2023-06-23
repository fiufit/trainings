package trainings

import (
	"context"
	"strconv"

	"github.com/fiufit/trainings/contracts/metrics"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type FavoriteAdder interface {
	AddToFavorite(ctx context.Context, userID string, trainingID uint) error
	RemoveFromFavorite(ctx context.Context, userID string, trainingID uint) error
}

type FavoriteAdderImpl struct {
	trainings repositories.TrainingPlans
	metrics   repositories.Metrics
	logger    *zap.Logger
}

func NewFavoriteAdderImpl(trainings repositories.TrainingPlans, metrics repositories.Metrics, logger *zap.Logger) FavoriteAdderImpl {
	return FavoriteAdderImpl{trainings: trainings, metrics: metrics, logger: logger}
}

func (uc *FavoriteAdderImpl) AddToFavorite(ctx context.Context, userID string, trainingID uint) error {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return err
	}
	err = uc.trainings.AddToFavorite(ctx, userID, trainingID, training.Version)
	if err != nil {
		return err
	}

	favoriteMetricReq := metrics.CreateMetricRequest{
		MetricType: "favorite_training",
		SubType:    strconv.Itoa(int(trainingID)),
	}

	uc.metrics.Create(ctx, favoriteMetricReq)
	return nil
}

func (uc *FavoriteAdderImpl) RemoveFromFavorite(ctx context.Context, userID string, trainingID uint) error {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return err
	}
	return uc.trainings.RemoveFromFavorite(ctx, userID, trainingID, training.Version)
}
