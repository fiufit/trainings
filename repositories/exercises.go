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

//go:generate mockery --name Exercises
type Exercises interface {
	CreateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error)
	DeleteExercise(ctx context.Context, exerciseID uint) error
	GetExerciseByID(ctx context.Context, exerciseID uint) (models.Exercise, error)
	UpdateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error)
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

func (repo ExerciseRepository) DeleteExercise(ctx context.Context, exerciseID uint) error {
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

func (repo ExerciseRepository) GetExerciseByID(ctx context.Context, exerciseID uint) (models.Exercise, error) {
	db := repo.db.WithContext(ctx)
	var exercise models.Exercise
	result := db.First(&exercise, "id = ?", exerciseID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Exercise{}, contracts.ErrExerciseNotFound
		}
		repo.logger.Error("Unable to get exercise", zap.Error(result.Error), zap.Uint("ID", exerciseID))
		return models.Exercise{}, result.Error
	}

	return exercise, nil
}

func (repo ExerciseRepository) UpdateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error) {
	db := repo.db.WithContext(ctx)
	result := db.Save(&exercise)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), contracts.ErrForeignKey.Error()) {
			return models.Exercise{}, contracts.ErrExerciseNotFound
		}
		repo.logger.Error("Unable to update exercise", zap.Error(result.Error), zap.Any("execise", exercise))
		return models.Exercise{}, result.Error
	}
	return exercise, nil
}
