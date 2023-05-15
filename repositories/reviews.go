package repositories

import (
	"context"
	"errors"
	"strings"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Reviews interface {
	CreateReview(ctx context.Context, review models.Review) (models.Review, error)
	//UpdateReview(ctx context.Context, review models.Review) (models.Review, error)
	//DeleteReview(ctx context.Context, review models.Review) error
	//GetReviewByID(ctx context.Context, reviewID uint) (models.Review, error)
}

type ReviewRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewReviewRepository(db *gorm.DB, logger *zap.Logger) ReviewRepository {
	return ReviewRepository{db: db, logger: logger}
}

func (repo ReviewRepository) CreateReview(ctx context.Context, review models.Review) (models.Review, error) {
	db := repo.db.WithContext(ctx)
	result := db.Create(&review)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), contracts.ErrForeignKey.Error()) {
			return models.Review{}, contracts.ErrTrainingPlanNotFound
		}
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return models.Review{}, contracts.ErrReviewAlreadyExists
		}
		repo.logger.Error("Unable to create training plan", zap.Error(result.Error), zap.Any("review", review))
		return models.Review{}, result.Error
	}
	return review, nil
}
