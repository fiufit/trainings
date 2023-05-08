package repositories

import (
	"context"

	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockery --name Exercises
type Exercises interface {
	CreateExercise(ctx context.Context, exercise models.Exercise) (models.Exercise, error)
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
