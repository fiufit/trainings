package repositories

import (
	"context"

	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Goals interface {
	Create(ctx context.Context, goal models.Goal) (models.Goal, error)
}

type GoalsRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewGoalsRepository(db *gorm.DB, logger *zap.Logger) GoalsRepository {
	return GoalsRepository{db: db, logger: logger}
}

func (repo GoalsRepository) Create(ctx context.Context, goal models.Goal) (models.Goal, error) {
	db := repo.db.WithContext(ctx)
	res := db.Create(&goal)
	if res.Error != nil {
		repo.logger.Error(res.Error.Error(), zap.Any("goal", goal))
		return models.Goal{}, res.Error
	}
	return goal, nil
}
