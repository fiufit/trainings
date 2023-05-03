package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/helpers"
	"github.com/fiufit/trainings/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"go.uber.org/zap/zaptest"
)

func TestCreateTrainingOk(t *testing.T) {

	ctx := context.Background()
	creationDate := time.Now()
	mpatch.PatchMethod(time.Now, func() time.Time {
		return creationDate
	})
	req := training.CreateTrainingRequest{
		Name:        "Test Name",
		Description: "Test Description",
		TrainerID:   "Test Trainer",
		Exercises:   []training.ExerciseRequest{},
	}
	trainingRepo := new(mocks.TrainingPlans)

	training := helpers.ConverToTrainingPlan(req, helpers.ConvertToExercises(req.Exercises))
	trainingRepo.On("CreateTrainingPlan", ctx, training).Return(training, nil)

	trainingUc := NewTrainingCreatorImpl(trainingRepo, zaptest.NewLogger(t))
	res, err := trainingUc.CreateTraining(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, res.TrainingPlan, training)
}
