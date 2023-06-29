package trainings

import (
	"context"
	"testing"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestNewTrainingDeleterImpl(t *testing.T) {
	trainingRepo := new(mocks.TrainingPlans)
	trainingUc := NewTrainingDeleterImpl(trainingRepo, zaptest.NewLogger(t))

	assert.NotNil(t, trainingUc)
	assert.NotNil(t, trainingUc.trainings)
	assert.NotNil(t, trainingUc.logger)
}

func TestDeleteTrainingOk(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)

	req := trainings.DeleteTrainingRequest{
		TrainerID:      trainerID,
		TrainingPlanID: trainingPlanID,
	}

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("DeleteTrainingPlan", ctx, req.TrainingPlanID).Return(nil)
	trainingRepo.On("GetTrainingByID", ctx, req.TrainingPlanID).Return(models.TrainingPlan{TrainerID: trainerID}, nil)

	trainingUc := NewTrainingDeleterImpl(trainingRepo, zaptest.NewLogger(t))
	err := trainingUc.DeleteTraining(ctx, req)

	assert.NoError(t, err)

}

func TestDeleteTrainingUnauthorizedTrainerError(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)

	req := trainings.DeleteTrainingRequest{
		TrainerID:      trainerID,
		TrainingPlanID: trainingPlanID,
	}

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("DeleteTrainingPlan", ctx, req.TrainingPlanID).Return(nil)
	trainingRepo.On("GetTrainingByID", ctx, req.TrainingPlanID).Return(models.TrainingPlan{TrainerID: "other"}, nil)

	trainingUc := NewTrainingDeleterImpl(trainingRepo, zaptest.NewLogger(t))
	err := trainingUc.DeleteTraining(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrUnauthorizedTrainer, err)
}

func TestDeleteTrainingNotFoundError(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)

	req := trainings.DeleteTrainingRequest{
		TrainerID:      trainerID,
		TrainingPlanID: trainingPlanID,
	}

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("GetTrainingByID", ctx, req.TrainingPlanID).Return(models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound)

	trainingUc := NewTrainingDeleterImpl(trainingRepo, zaptest.NewLogger(t))
	err := trainingUc.DeleteTraining(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrTrainingPlanNotFound, err)
}
