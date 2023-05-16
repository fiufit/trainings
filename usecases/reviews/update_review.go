package reviews

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ReviewUpdater interface {
	UpdateReview(ctx context.Context, req reviews.UpdateReviewRequest) (models.Review, error)
}

type ReviewUpdaterImpl struct {
	trainings repositories.TrainingPlans
	reviews   repositories.Reviews
	users     repositories.Users
	logger    *zap.Logger
}

func NewReviewUpdaterImpl(trainings repositories.TrainingPlans, reviews repositories.Reviews, users repositories.Users, logger *zap.Logger) ReviewUpdaterImpl {
	return ReviewUpdaterImpl{trainings: trainings, reviews: reviews, users: users, logger: logger}
}

func (uc *ReviewUpdaterImpl) UpdateReview(ctx context.Context, req reviews.UpdateReviewRequest) (models.Review, error) {
	_, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return models.Review{}, err
	}

	review, err := uc.reviews.GetReviewByID(ctx, req.ReviewID)
	if err != nil {
		return models.Review{}, err
	}

	if review.UserID != req.UserID {
		return models.Review{}, contracts.ErrUnauthorizedReviewer
	}

	patchedReview := uc.patchReviewModel(ctx, review, req)

	updatedReview, err := uc.reviews.UpdateReview(ctx, patchedReview)
	if err != nil {
		return models.Review{}, err
	}

	usr, _ := uc.users.GetUserByID(ctx, updatedReview.UserID)
	updatedReview.User = usr

	return updatedReview, nil
}

func (uc *ReviewUpdaterImpl) patchReviewModel(ctx context.Context, review models.Review, req reviews.UpdateReviewRequest) models.Review {
	if req.Score != 0 {
		review.Score = req.Score
	}

	if req.Comment != "" {
		review.Comment = req.Comment
	}

	return review
}
