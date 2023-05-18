package reviews

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ReviewCreator interface {
	CreateReview(ctx context.Context, req reviews.CreateReviewRequest) (models.Review, error)
}

type ReviewCreatorImpl struct {
	users     repositories.Users
	trainings repositories.TrainingPlans
	reviews   repositories.Reviews
	logger    *zap.Logger
}

func NewReviewCreatorImpl(trainings repositories.TrainingPlans, reviews repositories.Reviews, users repositories.Users, logger *zap.Logger) ReviewCreatorImpl {
	return ReviewCreatorImpl{trainings: trainings, reviews: reviews, users: users, logger: logger}
}

func (uc *ReviewCreatorImpl) CreateReview(ctx context.Context, req reviews.CreateReviewRequest) (models.Review, error) {
	usr, err := uc.users.GetUserByID(ctx, req.UserID)
	if err != nil {
		return models.Review{}, err
	}
	training, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return models.Review{}, err
	}
	if training.TrainerID == req.UserID {
		return models.Review{}, contracts.ErrSelfReview
	}
	newReview := models.Review{
		TrainingPlanID: req.TrainingPlanID,
		UserID:         req.UserID,
		User:           usr,
		Score:          req.Score,
		Comment:        req.Comment,
	}
	createdReview, err := uc.reviews.CreateReview(ctx, newReview)
	return createdReview, err
}
