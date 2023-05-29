package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockery --name TrainingPlans
type TrainingPlans interface {
	CreateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error)
	GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error)
	GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error)
	UpdateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error)
	DeleteTrainingPlan(ctx context.Context, trainingID uint) error
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

func (repo TrainingRepository) GetTrainingByID(ctx context.Context, trainingID uint) (models.TrainingPlan, error) {
	db := repo.db.WithContext(ctx)
	var training models.TrainingPlan
	result := db.
		Preload("Exercises").
		Preload("Reviews").
		Select("training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score").
		First(&training, "id = ?", trainingID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound
		}
		repo.logger.Error("Unable to get training plan", zap.Error(result.Error), zap.Uint("ID", trainingID))
		return models.TrainingPlan{}, result.Error
	}

	type Result struct {
		models.TrainingPlan
		MeanScore float32
	}

	var resultStruct Result
	result.Scan(&resultStruct)

	training.MeanScore = resultStruct.MeanScore

	return training, nil
}

func (repo TrainingRepository) GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	var res []models.TrainingPlan
	db := repo.db.WithContext(ctx)

	if req.Name != "" {
		likeName := fmt.Sprintf("%v%%", strings.ToLower(req.Name))
		db = db.Where("LOWER(name) LIKE ?", likeName)
	}

	if req.Description != "" {
		likeDescription := fmt.Sprintf("%%%v%%", req.Description)
		db = db.Where("LOWER(description) LIKE LOWER(?)", likeDescription)
	}

	if req.Difficulty != "" {
		db = db.Where("LOWER(difficulty) = ?", strings.ToLower(req.Difficulty))
	}

	if req.TrainerID != "" {
		db = db.Where("trainer_id = ?", req.TrainerID)
	}

	if req.MinDuration != 0 || req.MaxDuration != 0 {
		db = db.Where("duration >= ? AND (duration <= ? OR ? = 0)", req.MinDuration, req.MaxDuration, req.MaxDuration)
	}

	result := db.
		Scopes(database.Paginate(res, &req.Pagination, db)).
		Preload("Exercises").
		Preload("Reviews").
		Select("training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score").
		Order("mean_score DESC").
		Find(&res)

	if result.Error != nil {
		repo.logger.Error("Unable to get training plans with pagination", zap.Error(result.Error), zap.Any("request", req))
		return trainings.GetTrainingsResponse{}, result.Error
	}

	type Result struct {
		models.TrainingPlan
		MeanScore float32
	}

	var results []Result
	result.Scan(&results)

	for i := range res {
		res[i].MeanScore = results[i].MeanScore
	}

	return trainings.GetTrainingsResponse{TrainingPlans: res, Pagination: req.Pagination}, nil
}

func (repo TrainingRepository) UpdateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error) {
	db := repo.db.WithContext(ctx)
	result := db.Save(&training)
	if result.Error != nil {
		repo.logger.Error("Unable to update training plan", zap.Error(result.Error), zap.Any("training", training))
		return models.TrainingPlan{}, result.Error
	}
	return training, nil
}

func (repo TrainingRepository) DeleteTrainingPlan(ctx context.Context, trainingID uint) error {
	db := repo.db.WithContext(ctx)
	result := db.Select("Exercises", "Reviews").Delete(&models.TrainingPlan{ID: trainingID})
	if result.Error != nil {
		repo.logger.Error("Unable to delete training plan", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrTrainingPlanNotFound
	}
	return nil
}
