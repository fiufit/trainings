package trainings

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"go.uber.org/zap/zaptest"
)

func TestCreateTrainingOk(t *testing.T) {

	ctx := context.Background()
	creationDate := time.Now()
	_, err := mpatch.PatchMethod(time.Now, func() time.Time {
		return creationDate
	})
	if err != nil {
		t.Fatal(err)
	}
	req := trainings.CreateTrainingRequest{
		Name:        "Test Name",
		Description: "Test Description",
		TrainerID:   "Test Trainer",
		Exercises:   []trainings.ExerciseRequest{},
	}
	trainingRepo := new(mocks.TrainingPlans)
	userRepo := new(mocks.Users)

	training := trainings.ConverToTrainingPlan(req)
	trainingRepo.On("CreateTrainingPlan", ctx, training).Return(training, nil)
	userRepo.On("GetUserByID", ctx, req.TrainerID).Return(models.User{}, nil)

	trainingUc := NewTrainingCreatorImpl(trainingRepo, userRepo, zaptest.NewLogger(t))
	res, err := trainingUc.CreateTraining(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, res.TrainingPlan, training)
}

func TestCreateTrainingError(t *testing.T) {

	ctx := context.Background()
	req := trainings.CreateTrainingRequest{
		Name:        "Test Name",
		Description: "Test Description",
		TrainerID:   "Test Trainer",
		Exercises:   []trainings.ExerciseRequest{},
	}
	trainingRepo := new(mocks.TrainingPlans)
	userRepo := new(mocks.Users)

	training := trainings.ConverToTrainingPlan(req)
	trainingRepo.On("CreateTrainingPlan", ctx, training).Return(models.TrainingPlan{}, errors.New("repo error"))
	userRepo.On("GetUserByID", ctx, req.TrainerID).Return(models.User{}, nil)

	trainingUc := NewTrainingCreatorImpl(trainingRepo, userRepo, zaptest.NewLogger(t))
	res, err := trainingUc.CreateTraining(ctx, req)

	assert.Equal(t, res.TrainingPlan, models.TrainingPlan{})
	assert.Error(t, err)
}
