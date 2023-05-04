package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockery --name TrainingPlans
type TrainingPlans interface {
	CreateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error)
	GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error)
}

type TrainingRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTrainingRepository(db *gorm.DB, logger *zap.Logger) TrainingRepository {
	return TrainingRepository{db: db, logger: logger}
}

func (repo TrainingRepository) CreateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error) {
	db := repo.db.WithContext(ctx)
	result := db.Create(&training)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), contracts.ErrForeignKey.Error()) {
			return models.TrainingPlan{}, contracts.ErrUserNotFound
		}
		repo.logger.Error("Unable to create training plan", zap.Error(result.Error), zap.Any("training", training))
		return models.TrainingPlan{}, result.Error
	}
	return training, nil
}

func (repo TrainingRepository) GetTrainingPlans(ctx context.Context, req training.GetTrainingsRequest) (training.GetTrainingsResponse, error) {
	var res []models.TrainingPlan
	db := repo.db.WithContext(ctx)

	if req.Name != "" {
		likeName := fmt.Sprintf("%v%%", strings.ToLower(req.Name))
		db = db.Where("LOWER(name) LIKE ?", likeName).Preload("Exercises")
	}

	if req.Description != "" {
		likeDescription := fmt.Sprintf("%%%v%%", req.Description)
		db = db.Where("LOWER(description) LIKE LOWER(?)", likeDescription).Preload("Exercises")
	}

	if req.Difficulty != "" {
		db = db.Where("LOWER(difficulty) = ?", strings.ToLower(req.Difficulty)).Preload("Exercises")
	}

	if req.TrainerID != "" {
		db = db.Where("trainer_id = ?", req.TrainerID).Preload("Exercises")
	}

	if req.MinDuration != 0 || req.MaxDuration != 0 {
		db = db.Where("duration >= ? AND (duration <= ? OR ? = 0)", req.MinDuration, req.MaxDuration, req.MaxDuration).Preload("Exercises")
	}

	result := db.Scopes(database.Paginate(res, &req.Pagination, db)).Find(&res)
	if result.Error != nil {
		repo.logger.Error("Unable to get training plans with pagination", zap.Error(result.Error), zap.Any("request", req))
		return training.GetTrainingsResponse{}, result.Error
	}

	return training.GetTrainingsResponse{TrainingPlans: res, Pagination: req.Pagination}, nil
}
