package trainings

import (
	"context"
	"testing"
	"time"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestNewTrainingUpdaterImpl(t *testing.T) {
	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	trainingUc := NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, zaptest.NewLogger(t))

	assert.NotNil(t, trainingUc)
	assert.NotNil(t, trainingUc.trainings)
	assert.NotNil(t, trainingUc.firebase)
	assert.NotNil(t, trainingUc.logger)
}

func TestUpdateTrainingOk(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)
	currentTime := time.Now()

	req := trainings.UpdateTrainingRequest{
		ID: trainingPlanID,
		BaseTrainingRequest: trainings.BaseTrainingRequest{
			TrainerID: trainerID,
			Name:      "updated name",
			Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
			Exercises: []trainings.ExerciseRequest{
				{
					Title:       "exercise1",
					Description: "description1",
				},
			},
		},
	}

	oldTraining := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: trainerID,
		Name:      "old name",
		CreatedAt: currentTime,
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	updatedTraining := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: trainerID,
		Name:      req.Name,
		CreatedAt: currentTime,
		Tags:      req.Tags,
		Exercises: []models.Exercise{
			{
				Title:       req.Exercises[0].Title,
				Description: req.Exercises[0].Description,
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)

	trainingRepo.On("UpdateTrainingPlan", ctx, updatedTraining).Return(updatedTraining, nil)
	trainingRepo.On("GetTrainingByID", ctx, req.ID).Return(oldTraining, nil)
	firebaseRepo.On("FillTrainingPicture", ctx, &updatedTraining).Return(nil)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, zaptest.NewLogger(t))
	updated, err := trainingUc.UpdateTrainingPlan(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, updatedTraining, updated)
}

func TestUpdateTrainingFailsWithRepositoryGetTrainingByIDErr(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)

	req := trainings.UpdateTrainingRequest{
		ID: trainingPlanID,
		BaseTrainingRequest: trainings.BaseTrainingRequest{
			TrainerID: trainerID,
			Name:      "updated name",
			Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
			Exercises: []trainings.ExerciseRequest{
				{
					Title:       "exercise1",
					Description: "description1",
				},
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)

	trainingRepo.On("GetTrainingByID", ctx, req.ID).Return(models.TrainingPlan{}, contracts.ErrInternal)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, zaptest.NewLogger(t))
	_, err := trainingUc.UpdateTrainingPlan(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrInternal, err)
}

func TestUpdateTrainingFailsWithRepositoryUpdateTrainingPlanErr(t *testing.T) {
	ctx := context.Background()
	currentTime := time.Now()

	trainingPlan := models.TrainingPlan{
		ID:        1,
		TrainerID: "test",
		Name:      "updated name",
		CreatedAt: currentTime,
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	req := trainings.UpdateTrainingRequest{
		ID: trainingPlan.ID,
		BaseTrainingRequest: trainings.BaseTrainingRequest{
			TrainerID: trainingPlan.TrainerID,
			Name:      trainingPlan.Name,
			Exercises: []trainings.ExerciseRequest{
				{
					Title:       trainingPlan.Exercises[0].Title,
					Description: trainingPlan.Exercises[0].Description,
				},
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)

	trainingRepo.On("GetTrainingByID", ctx, req.ID).Return(trainingPlan, nil)
	trainingRepo.On("UpdateTrainingPlan", ctx, trainingPlan).Return(models.TrainingPlan{}, contracts.ErrInternal)

	firebaseRepo.On("FillTrainingPicture", ctx, &trainingPlan).Return(nil)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, zaptest.NewLogger(t))
	_, err := trainingUc.UpdateTrainingPlan(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrInternal, err)

}

func TestUpdateTrainingUnauthorizedTrainerErr(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)
	currentTime := time.Now()

	req := trainings.UpdateTrainingRequest{
		ID: trainingPlanID,
		BaseTrainingRequest: trainings.BaseTrainingRequest{
			TrainerID: trainerID,
			Name:      "updated name",
			Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
			Exercises: []trainings.ExerciseRequest{
				{
					Title:       "exercise1",
					Description: "description1",
				},
			},
		},
	}

	oldTraining := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: "other",
		Name:      "old name",
		CreatedAt: currentTime,
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("GetTrainingByID", ctx, req.ID).Return(oldTraining, nil)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, nil, zaptest.NewLogger(t))
	_, err := trainingUc.UpdateTrainingPlan(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrUnauthorizedTrainer, err)
}

func TestDisableTrainingPlanOk(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("UpdateDisabledStatus", ctx, trainingID, true).Return(nil)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, nil, zaptest.NewLogger(t))
	err := trainingUc.DisableTrainingPlan(ctx, trainingID)

	assert.NoError(t, err)
}

func TestDisableTrainingPlanFailsWithRepositoryUpdateDisabledStatusErr(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("UpdateDisabledStatus", ctx, trainingID, true).Return(contracts.ErrInternal)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, nil, zaptest.NewLogger(t))
	err := trainingUc.DisableTrainingPlan(ctx, trainingID)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrInternal, err)
}

func TestEnableTrainingPlanOk(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("UpdateDisabledStatus", ctx, trainingID, false).Return(nil)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, nil, zaptest.NewLogger(t))
	err := trainingUc.EnableTrainingPlan(ctx, trainingID)

	assert.NoError(t, err)
}

func TestEnableTrainingPlanFailsWithRepositoryUpdateDisabledStatusErr(t *testing.T) {
	ctx := context.Background()
	trainingID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)

	trainingRepo.On("UpdateDisabledStatus", ctx, trainingID, false).Return(contracts.ErrInternal)

	trainingUc := NewTrainingUpdaterImpl(trainingRepo, nil, zaptest.NewLogger(t))
	err := trainingUc.EnableTrainingPlan(ctx, trainingID)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrInternal, err)
}
