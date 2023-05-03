package repositories

import (
	"context"
	"strings"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TrainingPlans interface {
	CreateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error)
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
