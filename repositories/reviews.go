package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Reviews interface {
	CreateReview(ctx context.Context, review models.Review) (models.Review, error)
	UpdateReview(ctx context.Context, review models.Review) (models.Review, error)
	DeleteReview(ctx context.Context, reviewID uint) error
	GetReviewByID(ctx context.Context, reviewID uint) (models.Review, error)
	GetReviews(ctx context.Context, req reviews.GetReviewsRequest) (reviews.GetReviewsResponse, error)
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
		repo.logger.Error("Unable to create review", zap.Error(result.Error), zap.Any("review", review))
		return models.Review{}, result.Error
	}
	return review, nil
}

func (repo ReviewRepository) UpdateReview(ctx context.Context, review models.Review) (models.Review, error) {
	db := repo.db.WithContext(ctx)
	result := db.Save(&review)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Review{}, contracts.ErrReviewNotFound
		}
		repo.logger.Error("Unable to update review", zap.Error(result.Error), zap.Any("review", review))
		return models.Review{}, result.Error
	}
	return review, nil
}

func (repo ReviewRepository) DeleteReview(ctx context.Context, reviewID uint) error {
	db := repo.db.WithContext(ctx)
	var review models.Review
	result := db.Delete(&review, "id = ?", reviewID)
	if result.Error != nil {
		repo.logger.Error("Unable to delete review", zap.Error(result.Error), zap.Any("review", review))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrReviewNotFound
	}
	return nil
}

func (repo ReviewRepository) GetReviewByID(ctx context.Context, reviewID uint) (models.Review, error) {
	db := repo.db.WithContext(ctx)
	var review models.Review
	result := db.First(&review, "id = ?", reviewID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Review{}, contracts.ErrReviewNotFound
		}
		repo.logger.Error("Unable to get review", zap.Error(result.Error), zap.Uint("ID", reviewID))
		return models.Review{}, result.Error
	}
	return review, nil
}

func (repo ReviewRepository) GetReviews(ctx context.Context, req reviews.GetReviewsRequest) (reviews.GetReviewsResponse, error) {
	var res []models.Review
	db := repo.db.WithContext(ctx)
	db = db.Where("training_plan_id = ?", req.TrainingPlanID)

	if req.MinScore != 0 || req.MaxScore != 0 {
		db = db.Where("score >= ? AND (score <= ? OR ? = 0)", req.MinScore, req.MaxScore, req.MaxScore)
	}

	if req.UserID != "" {
		db = db.Where("user_id = ?", req.UserID)
	}

	if req.Comment != "" {
		likeComment := fmt.Sprintf("%%%v%%", req.Comment)
		db = db.Where("LOWER(comment) LIKE LOWER(?)", likeComment)
	}

	result := db.Scopes(database.Paginate(res, &req.Pagination, db)).Find(&res)

	if result.Error != nil {
		repo.logger.Error("Unable to get reviews with pagination", zap.Error(result.Error), zap.Any("request", req))
		return reviews.GetReviewsResponse{}, result.Error
	}

	return reviews.GetReviewsResponse{Reviews: res, Pagination: req.Pagination}, nil
}
