package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

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
	UpdateBySession(ctx context.Context, session models.TrainingSession) error
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

func (repo GoalsRepository) UpdateBySession(ctx context.Context, session models.TrainingSession) error {
	db := repo.db.WithContext(ctx).Where("deadline > ? AND user_id = ? AND goal_value > goal_value_progress", time.Now(), session.UserID)

	tagStrings := make([]string, len(session.TrainingPlan.Tags))
	for i, tag := range session.TrainingPlan.Tags {
		tagStrings[i] = tag.Name
	}

	if session.TrainingPlan.Difficulty != "" && session.TrainingPlan.Tags != nil {
		db.Model(&models.Goal{}).
			Where("goal_type = ?", "sessions count").
			Where("goal_subtype = ? OR goal_subtype IN (?)", strings.ToLower(session.TrainingPlan.Difficulty), tagStrings).
			UpdateColumn("goal_value_progress", gorm.Expr("goal_value_progress + ?", 1))

		if db.Error != nil {
			fmt.Println("Unable to update sessions count goals")
			repo.logger.Error("Unable to update sessions count goals", zap.Error(db.Error))
			return db.Error
		}
	}

	if session.SecondsCount > 0 {
		db.Model(&models.Goal{}).
			Where("goal_type = ?", "minutes count").
			UpdateColumn("goal_value_progress", gorm.Expr("goal_value_progress + ?", session.SecondsCount/60))

		if db.Error != nil {
			fmt.Println("Unable to update minutes count goals")
			repo.logger.Error("Unable to update minutes count goals", zap.Error(db.Error))
			return db.Error
		}
	}

	if session.StepCount > 0 {
		db.Model(&models.Goal{}).
			Where("goal_type = ?", "step count").
			UpdateColumn("goal_value_progress", gorm.Expr("goal_value_progress + ?", session.StepCount))

		if db.Error != nil {
			fmt.Println("Unable to update step count goals")
			repo.logger.Error("Unable to update step count goals", zap.Error(db.Error))
			return db.Error
		}
	}

	return nil

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
