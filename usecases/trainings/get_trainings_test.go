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

func TestNewTrainingGetterImpl(t *testing.T) {
	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	assert.NotNil(t, trainingUc)
	assert.NotNil(t, trainingUc.trainings)
	assert.NotNil(t, trainingUc.firebase)
	assert.NotNil(t, trainingUc.users)
	assert.NotNil(t, trainingUc.logger)
}

func TestGetTrainingByIdOk(t *testing.T) {
	ctx := context.Background()
	trainingPlanID := uint(1)

	training := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: "test",
		Name:      "test",
		CreatedAt: time.Now(),
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingRepo.On("GetTrainingByID", ctx, trainingPlanID).Return(training, nil)
	firebaseRepo.On("FillTrainingPicture", ctx, &training).Return(nil)

	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetTrainingByID(ctx, trainingPlanID)

	assert.NoError(t, err)
	assert.Equal(t, trainingPlanID, resp.ID)
	assert.Equal(t, training.TrainerID, resp.TrainerID)
	assert.Equal(t, training.Name, resp.Name)
	assert.Equal(t, training.Tags, resp.Tags)
	assert.Equal(t, training.Exercises, resp.Exercises)

}

func TestGetTrainingByIdNotFoundErr(t *testing.T) {
	ctx := context.Background()
	trainingPlanID := uint(1)

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingRepo.On("GetTrainingByID", ctx, trainingPlanID).Return(models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound)

	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetTrainingByID(ctx, trainingPlanID)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrTrainingPlanNotFound, err)
	assert.Equal(t, models.TrainingPlan{}, resp)
}

func TestGetTrainingPlansOk(t *testing.T) {
	ctx := context.Background()
	trainerID := "test"
	trainingPlanID := uint(1)

	req := trainings.GetTrainingsRequest{
		TrainerID: trainerID,
	}

	training := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: trainerID,
		Name:      "test",
		CreatedAt: time.Now(),
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingRepo.On("GetTrainingPlans", ctx, req).Return(trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, nil)
	firebaseRepo.On("FillTrainingPicture", ctx, &training).Return(nil)

	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetTrainingPlans(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, resp)
}

func TestGetRecommendedPlansOk(t *testing.T) {
	ctx := context.Background()
	userID := "user_test"
	trainerID := "trainer_test"
	trainingPlanID := uint(1)

	user := models.User{
		ID:        userID,
		Interests: []models.UserInterest{{Name: "speed"}, {Name: "strength"}},
	}

	ucReq := trainings.GetTrainingsRequest{
		UserID: userID,
	}

	cleanReq := trainings.GetTrainingsRequest{
		TagStrings: []string{"speed", "strength"},
		Tags:       []models.Tag{{Name: "speed"}, {Name: "strength"}},
	}

	training := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: trainerID,
		Name:      "test",
		CreatedAt: time.Now(),
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingRepo.On("GetTrainingPlans", ctx, cleanReq).Return(trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, nil)
	firebaseRepo.On("FillTrainingPicture", ctx, &training).Return(nil)
	userRepo.On("GetUserByID", ctx, userID).Return(user, nil)

	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetRecommendedPlans(ctx, ucReq)

	assert.NoError(t, err)
	assert.Equal(t, trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, resp)
}

func TestGetRecommendedPlansForNonExistingUser(t *testing.T) {
	ctx := context.Background()
	userID := "user_test"

	ucReq := trainings.GetTrainingsRequest{
		UserID: userID,
	}

	userRepo := new(mocks.Users)
	userRepo.On("GetUserByID", ctx, userID).Return(models.User{}, contracts.ErrUserNotFound)

	trainingUc := NewTrainingGetterImpl(nil, nil, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetRecommendedPlans(ctx, ucReq)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrUserNotFound, err)
	assert.Equal(t, trainings.GetTrainingsResponse{}, resp)
}

func TestGetFavoritePlansOk(t *testing.T) {
	ctx := context.Background()
	userID := "user_test"
	trainerID := "trainer_test"
	trainingPlanID := uint(1)

	user := models.User{
		ID:        userID,
		Interests: []models.UserInterest{{Name: "speed"}, {Name: "strength"}},
	}

	ucReq := trainings.GetFavoritesRequest{
		UserID: userID,
	}

	training := models.TrainingPlan{
		ID:        trainingPlanID,
		TrainerID: trainerID,
		Name:      "test",
		CreatedAt: time.Now(),
		Tags:      []models.Tag{{Name: "tag1"}, {Name: "tag2"}},
		Exercises: []models.Exercise{
			{
				Title:       "exercise1",
				Description: "description1",
			},
		},
	}

	trainingRepo := new(mocks.TrainingPlans)
	firebaseRepo := new(mocks.Firebase)
	userRepo := new(mocks.Users)
	trainingRepo.On("GetFavoriteTrainings", ctx, trainings.GetFavoritesRequest{UserID: userID}).Return(trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, nil)
	firebaseRepo.On("FillTrainingPicture", ctx, &training).Return(nil)
	userRepo.On("GetUserByID", ctx, userID).Return(user, nil)

	trainingUc := NewTrainingGetterImpl(trainingRepo, firebaseRepo, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetFavoritePlans(ctx, ucReq)

	assert.NoError(t, err)
	assert.Equal(t, trainings.GetTrainingsResponse{TrainingPlans: []models.TrainingPlan{training}}, resp)
}

func TestGetFavoritePlansForNonExistingUserErr(t *testing.T) {
	ctx := context.Background()
	userID := "user_test"

	ucReq := trainings.GetFavoritesRequest{
		UserID: userID,
	}

	userRepo := new(mocks.Users)
	userRepo.On("GetUserByID", ctx, userID).Return(models.User{}, contracts.ErrUserNotFound)

	trainingUc := NewTrainingGetterImpl(nil, nil, userRepo, zaptest.NewLogger(t))

	resp, err := trainingUc.GetFavoritePlans(ctx, ucReq)

	assert.Error(t, err)
	assert.Equal(t, contracts.ErrUserNotFound, err)
	assert.Equal(t, trainings.GetTrainingsResponse{}, resp)
}
