package repositories

import (
	"context"
	"errors"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Goals interface {
	Create(ctx context.Context, goal models.Goal) (models.Goal, error)
	GetByID(ctx context.Context, goalID uint) (models.Goal, error)
	Get(ctx context.Context, req goals.GetGoalsRequest) ([]models.Goal, error)
	Update(ctx context.Context, goal models.Goal) (models.Goal, error)
	//UpdateBySession(ctx context.Context)
	Delete(ctx context.Context, goalID uint) error
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

func (repo GoalsRepository) GetByID(ctx context.Context, goalID uint) (models.Goal, error) {
	db := repo.db.WithContext(ctx)
	var goal models.Goal
	result := db.First(&goal, "id = ?", goalID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Goal{}, contracts.ErrGoalNotFound
		}
		repo.logger.Error("Unable to get goal", zap.Error(result.Error), zap.Uint("ID", goalID))
		return models.Goal{}, result.Error
	}
	return goal, nil
}

func (repo GoalsRepository) Get(ctx context.Context, req goals.GetGoalsRequest) ([]models.Goal, error) {
	var goals []models.Goal
	db := repo.db.WithContext(ctx)

	if req.UserID != "" {
		db = db.Where("user_id = ?", req.UserID)
	}

	if req.GoalType != "" {
		db = db.Where("goal_type = LOWER(?)", req.GoalType)
	}

	if req.GoalSubtype != "" {
		db = db.Where("goal_subtype = LOWER(?)", req.GoalSubtype)
	}

	if !req.Deadline.IsZero() {
		db = db.Where("deadline < ?", req.Deadline)
	}
	res := db.Order("created_at desc").Find(&goals)

	if res.Error != nil {
		repo.logger.Error(res.Error.Error(), zap.Any("req", req))
		return nil, res.Error
	}
	return goals, nil

}

func (repo GoalsRepository) Update(ctx context.Context, goal models.Goal) (models.Goal, error) {
	db := repo.db.WithContext(ctx)
	result := db.Save(&goal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Goal{}, contracts.ErrGoalNotFound
		}
		repo.logger.Error("Unable to update goal", zap.Error(result.Error), zap.Any("goal", goal))
		return models.Goal{}, result.Error
	}
	return goal, nil
}

func (repo GoalsRepository) Delete(ctx context.Context, goalID uint) error {
	db := repo.db.WithContext(ctx)
	var goal models.Goal
	result := db.Delete(&goal, "id = ?", goalID)
	if result.Error != nil {
		repo.logger.Error("Unable to delete goal", zap.Error(result.Error), zap.Any("goal", goal))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrGoalNotFound
	}
	return nil
}
