package reviews

import (
	"context"

	"github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ReviewGetter interface {
	GetReviews(ctx context.Context, req reviews.GetReviewsRequest) (reviews.GetReviewsResponse, error)
	GetReviewByID(ctx context.Context, trainingPlanID uint, reviewID uint) (models.Review, error)
}

type ReviewGetterImpl struct {
	trainings repositories.TrainingPlans
	reviews   repositories.Reviews
	users     repositories.Users
	logger    *zap.Logger
}

func NewReviewGetterImpl(trainings repositories.TrainingPlans, reviews repositories.Reviews, users repositories.Users, logger *zap.Logger) ReviewGetterImpl {
	return ReviewGetterImpl{trainings: trainings, reviews: reviews, users: users, logger: logger}
}

func (uc *ReviewGetterImpl) GetReviews(ctx context.Context, req reviews.GetReviewsRequest) (reviews.GetReviewsResponse, error) {
	_, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return reviews.GetReviewsResponse{}, err
	}
	res, err := uc.reviews.GetReviews(ctx, req)
	if err != nil {
		return reviews.GetReviewsResponse{}, err
	}
	for i := range res.Reviews {
		usr, _ := uc.users.GetUserByID(ctx, res.Reviews[i].UserID)
		res.Reviews[i].User = usr
	}
	return res, nil
}

func (uc *ReviewGetterImpl) GetReviewByID(ctx context.Context, trainingPlanID uint, reviewID uint) (models.Review, error) {
	_, err := uc.trainings.GetTrainingByID(ctx, trainingPlanID)
	if err != nil {
		return models.Review{}, err
	}
	review, err := uc.reviews.GetReviewByID(ctx, reviewID)
	if err != nil {
		return models.Review{}, err
	}
	usr, _ := uc.users.GetUserByID(ctx, review.UserID)
	review.User = usr
	return review, nil
}
