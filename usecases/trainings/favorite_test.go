package trainings

import (
	"context"
	"strconv"
	"testing"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/metrics"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestNewFavoriteAdderImpl(t *testing.T) {
	trainingRepo := new(mocks.TrainingPlans)
	metricsRepo := new(mocks.Metrics)

	trainingUc := NewFavoriteAdderImpl(trainingRepo, metricsRepo, zaptest.NewLogger(t))

	assert.NotNil(t, trainingUc)
	assert.NotNil(t, trainingUc.trainings)
	assert.NotNil(t, trainingUc.metrics)
	assert.NotNil(t, trainingUc.logger)
}

func TestAddToFavoriteOk(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)
	trainingVersion := uint(2)
	userID := "user_id"

	trainingRepo := new(mocks.TrainingPlans)
	metricsRepo := new(mocks.Metrics)

	trainingRepo.On("GetTrainingByID", ctx, trainingID).Return(models.TrainingPlan{Version: trainingVersion}, nil)
	trainingRepo.On("AddToFavorite", ctx, userID, trainingID, trainingVersion).Return(nil)
	metricsRepo.On("Create", ctx, metrics.CreateMetricRequest{MetricType: "favorite_training", SubType: strconv.Itoa(int(trainingID))}).Return(nil)

	trainingUc := NewFavoriteAdderImpl(trainingRepo, metricsRepo, zaptest.NewLogger(t))

	err := trainingUc.AddToFavorite(ctx, userID, trainingID)
	assert.NoError(t, err)
}

func TestAddToFavoriteForNonExistingTraningErr(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)
	metricsRepo := new(mocks.Metrics)

	trainingRepo.On("GetTrainingByID", ctx, trainingID).Return(models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound)

	trainingUc := NewFavoriteAdderImpl(trainingRepo, metricsRepo, zaptest.NewLogger(t))

	err := trainingUc.AddToFavorite(ctx, "test", trainingID)
	assert.Error(t, err)
}

func TestRemoveFromFavoriteOk(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)
	trainingVersion := uint(2)
	userID := "user_id"

	trainingRepo := new(mocks.TrainingPlans)
	metricsRepo := new(mocks.Metrics)

	trainingRepo.On("GetTrainingByID", ctx, trainingID).Return(models.TrainingPlan{Version: trainingVersion}, nil)
	trainingRepo.On("RemoveFromFavorite", ctx, userID, trainingID, trainingVersion).Return(nil)

	trainingUc := NewFavoriteAdderImpl(trainingRepo, metricsRepo, zaptest.NewLogger(t))

	err := trainingUc.RemoveFromFavorite(ctx, userID, trainingID)
	assert.NoError(t, err)
}

func TestRemoveFromFavoriteForNonExistigTrainingErr(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)
	metricsRepo := new(mocks.Metrics)

	trainingRepo.On("GetTrainingByID", ctx, trainingID).Return(models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound)

	trainingUc := NewFavoriteAdderImpl(trainingRepo, metricsRepo, zaptest.NewLogger(t))

	err := trainingUc.RemoveFromFavorite(ctx, "test", trainingID)
	assert.Error(t, err)
}
