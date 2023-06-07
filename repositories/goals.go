package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/fiufit/trainings/contracts"
	gContracts "github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Goals interface {
	Create(ctx context.Context, goal models.Goal) (models.Goal, error)
	GetByID(ctx context.Context, goalID uint) (models.Goal, error)
	Get(ctx context.Context, req gContracts.GetGoalsRequest) ([]models.Goal, error)
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

func (repo GoalsRepository) Get(ctx context.Context, req gContracts.GetGoalsRequest) ([]models.Goal, error) {
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
	db := repo.db.WithContext(ctx).Where("", time.Now(), session.UserID)

	tagStrings := make([]string, len(session.TrainingPlan.Tags))
	for i, tag := range session.TrainingPlan.Tags {
		tagStrings[i] = tag.Name
	}

	if session.TrainingPlan.Difficulty != "" && session.TrainingPlan.Tags != nil {

		res := db.Exec("UPDATE goals SET goal_value_progress = goal_value_progress + 1 WHERE goal_type = 'sessions count' "+
			"AND (goal_subtype = ? OR goal_subtype IN (?)) AND deadline > NOW() AND user_id = ? AND goal_value > goal_value_progress",
			strings.ToLower(session.TrainingPlan.Difficulty),
			tagStrings,
			session.UserID,
		)

		if res.Error != nil {
			repo.logger.Error("Unable to update sessions count goals", zap.Error(res.Error))
			return res.Error
		}
	}

	if session.SecondsCount > 0 {

		res := db.Exec("UPDATE goals SET goal_value_progress = LEAST(goal_value_progress + ?, goal_value) WHERE goal_type = 'minutes count' "+
			"AND deadline > NOW() AND user_id = ? AND goal_value > goal_value_progress",
			int(session.SecondsCount/60),
			session.UserID,
		)

		if res.Error != nil {
			repo.logger.Error("Unable to update minutes count goals", zap.Error(res.Error))
			return res.Error
		}
	}

	if session.StepCount > 0 {
		res := db.Exec("UPDATE goals SET goal_value_progress = LEAST(goal_value_progress + ?, goal_value) WHERE goal_type = 'step count' "+
			"AND deadline > NOW() AND user_id = ? AND goal_value > goal_value_progress",
			session.StepCount,
			session.UserID,
		)

		if res.Error != nil {
			repo.logger.Error("Unable to update step count goals", zap.Error(res.Error))
			return res.Error
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
