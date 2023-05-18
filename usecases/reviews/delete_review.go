package reviews

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type ReviewDeleter interface {
	DeleteReview(ctx context.Context, req reviews.DeleteReviewRequest) error
}

type ReviewDeleterImpl struct {
	trainings repositories.TrainingPlans
	reviews   repositories.Reviews
	logger    *zap.Logger
}

func NewReviewDeleterImpl(trainings repositories.TrainingPlans, reviews repositories.Reviews, logger *zap.Logger) ReviewDeleterImpl {
	return ReviewDeleterImpl{trainings: trainings, reviews: reviews, logger: logger}
}

func (uc *ReviewDeleterImpl) DeleteReview(ctx context.Context, req reviews.DeleteReviewRequest) error {
	_, err := uc.trainings.GetTrainingByID(ctx, req.TrainingPlanID)
	if err != nil {
		return err
	}

	review, err := uc.reviews.GetReviewByID(ctx, req.ReviewID)
	if err != nil {
		return err
	}

	if review.UserID != req.UserID {
		return contracts.ErrUnauthorizedReviewer
	}
	return uc.reviews.DeleteReview(ctx, req.ReviewID)
}
