package trainings

import (
	"context"

	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingGetter interface {
	GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error)
	GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error)
	GetRecommendedPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error)
	GetFavoritePlans(ctx context.Context, req trainings.GetFavoritesRequest) (trainings.GetTrainingsResponse, error)
}

type TrainingGetterImpl struct {
	trainings repositories.TrainingPlans
	firebase  repositories.Firebase
	users     repositories.Users
	logger    *zap.Logger
}

func NewTrainingGetterImpl(trainings repositories.TrainingPlans, firebase repositories.Firebase, users repositories.Users, logger *zap.Logger) TrainingGetterImpl {
	return TrainingGetterImpl{trainings: trainings, firebase: firebase, users: users, logger: logger}
}

func (uc *TrainingGetterImpl) GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	res, err := uc.trainings.GetTrainingPlans(ctx, req)
	if err != nil {
		return res, err
	}
	for i := range res.TrainingPlans {
		uc.firebase.FillTrainingPicture(ctx, &res.TrainingPlans[i])
	}
	return res, nil
}

func (uc *TrainingGetterImpl) GetRecommendedPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	user, err := uc.users.GetUserByID(ctx, req.UserID)
	if err != nil {
		return trainings.GetTrainingsResponse{}, err
	}

	// ignore every other parameter in the request except pagination if the userID is set
	cleanReq := trainings.GetTrainingsRequest{Pagination: req.Pagination}

	interestStrings := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interestStrings[i] = interest.Name
	}
	cleanReq.TagStrings = interestStrings
	if err = cleanReq.Validate(); err != nil {
		return trainings.GetTrainingsResponse{}, err
	}

	return uc.GetTrainingPlans(ctx, cleanReq)
}

func (uc *TrainingGetterImpl) GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error) {
	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return training, err
	}
	uc.firebase.FillTrainingPicture(ctx, &training)
	return training, nil
}

func (uc *TrainingGetterImpl) GetFavoritePlans(ctx context.Context, req trainings.GetFavoritesRequest) (trainings.GetTrainingsResponse, error) {
	_, err := uc.users.GetUserByID(ctx, req.UserID)
	if err != nil {
		return trainings.GetTrainingsResponse{}, err
	}
	res, err := uc.trainings.GetFavoriteTrainings(ctx, req)
	if err != nil {
		return res, err
	}
	for i := range res.TrainingPlans {
		uc.firebase.FillTrainingPicture(ctx, &res.TrainingPlans[i])
	}
	return res, nil
}
