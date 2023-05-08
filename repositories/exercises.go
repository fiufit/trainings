package repositories

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockery --name Exercises
type Exercises interface {
	CreateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error)
	DeleteExercise(ctx context.Context, exerciseID string) error
}

type ExerciseRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewExerciseRepository(db *gorm.DB, logger *zap.Logger) ExerciseRepository {
	return ExerciseRepository{db: db, logger: logger}
}

func (repo ExerciseRepository) CreateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error) {
	db := repo.db.WithContext(ctx)
	result := db.Create(&exercise)
	if result.Error != nil {
		repo.logger.Error("Unable to create training plan", zap.Error(result.Error), zap.Any("exercise", exercise))
		return models.Exercise{}, result.Error
	}
	return exercise, nil
}

func (repo ExerciseRepository) DeleteExercise(ctx context.Context, exerciseID string) error {
	db := repo.db.WithContext(ctx)
	var exercise models.Exercise
	result := db.Delete(&exercise, "id = ?", exerciseID)
	if result.Error != nil {
		repo.logger.Error("Unable to delete exercise", zap.Error(result.Error), zap.Any("exercise", exercise))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrExerciseNotFound
	}
	return nil
}
