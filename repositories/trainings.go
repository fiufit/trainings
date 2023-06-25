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
	AddToFavorite(ctx context.Context, userID string, trainingID uint, trainingVersion uint) error
	RemoveFromFavorite(ctx context.Context, userID string, trainingID uint, trainingVersion uint) error
	GetFavoriteTrainings(ctx context.Context, req trainings.GetFavoritesRequest) (trainings.GetTrainingsResponse, error)
	UpdateDisabledStatus(ctx context.Context, trainingID uint, disabled bool) error
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
	training.Version = 1
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
		Preload("Tags").
		Where("disabled = false").
		Select(`training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score,
					(SELECT COUNT(*) FROM favorites WHERE favorites.training_plan_id = training_plans.id) AS favorites_count, 
					(SELECT COUNT(*) FROM training_sessions WHERE training_sessions.training_plan_id = training_plans.id) as sessions_count`).
		First(&training, "id = ?", trainingID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound
		}
		repo.logger.Error("Unable to get training plan", zap.Error(result.Error), zap.Uint("ID", trainingID))
		return models.TrainingPlan{}, result.Error
	}

	var resultStruct Result
	result.Scan(&resultStruct)

	training.MeanScore = resultStruct.MeanScore
	training.FavoritesCount = resultStruct.FavoritesCount
	training.SessionsCount = resultStruct.SessionsCount

	return training, nil
}

func (repo TrainingRepository) GetTrainingByIDAndVersion(ctx context.Context, trainingID uint, version uint) (models.TrainingPlan, error) {
	db := repo.db.WithContext(ctx)
	var training models.TrainingPlan
	result := db.Unscoped().
		Preload("Exercises").
		Preload("Reviews").
		Preload("Tags").
		Where("disabled = false").
		Select(`training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score,
					(SELECT COUNT(*) FROM favorites WHERE favorites.training_plan_id = training_plans.id) AS favorites_count, 
					(SELECT COUNT(*) FROM training_sessions WHERE training_sessions.training_plan_id = training_plans.id) as sessions_count`).
		First(&training, "id = ? AND version = ?", trainingID, version)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.TrainingPlan{}, contracts.ErrTrainingPlanNotFound
		}
		repo.logger.Error("Unable to get training plan", zap.Error(result.Error), zap.Uint("ID", trainingID))
		return models.TrainingPlan{}, result.Error
	}

	var resultStruct Result
	result.Scan(&resultStruct)

	training.MeanScore = resultStruct.MeanScore
	training.FavoritesCount = resultStruct.FavoritesCount
	training.SessionsCount = resultStruct.SessionsCount

	return training, nil
}
func (repo TrainingRepository) GetTrainingPlans(ctx context.Context, req trainings.GetTrainingsRequest) (trainings.GetTrainingsResponse, error) {
	var res []models.TrainingPlan
	db := repo.db.WithContext(ctx)

	if req.Disabled == nil {
		db = db.Where("disabled = ?", false)
	} else {
		db = db.Where("disabled = ?", *req.Disabled)
	}

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
	if len(req.Tags) > 0 {
		db = db.InnerJoins("INNER JOIN training_plan_tags ON training_plan_tags.training_plan_id = training_plans.id AND "+
			"training_plan_tags.training_plan_version = training_plans.version").Distinct().Where("training_plan_tags.tag_name IN (?)", req.TagStrings)
	}

	db = db.Select(`training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score,
					(SELECT COUNT(*) FROM favorites WHERE favorites.training_plan_id = training_plans.id) AS favorites_count, 
					(SELECT COUNT(*) FROM training_sessions WHERE training_sessions.training_plan_id = training_plans.id) as sessions_count`).
		Order("mean_score DESC")

	result := db.
		Scopes(database.Paginate(res, &req.Pagination, db)).
		Preload("Exercises").
		Preload("Reviews").
		Preload("Tags").
		Find(&res)

	if result.Error != nil {
		repo.logger.Error("Unable to get training plans with pagination", zap.Error(result.Error), zap.Any("request", req))
		return trainings.GetTrainingsResponse{}, result.Error
	}

	var results []Result
	result.Scan(&results)

	for i := range res {
		res[i].MeanScore = results[i].MeanScore
		res[i].FavoritesCount = results[i].FavoritesCount
		res[i].SessionsCount = results[i].SessionsCount
	}

	return trainings.GetTrainingsResponse{TrainingPlans: res, Pagination: req.Pagination}, nil
}

func (repo TrainingRepository) UpdateTrainingPlan(ctx context.Context, training models.TrainingPlan) (models.TrainingPlan, error) {
	db := repo.db.WithContext(ctx)

	oldPlan, err := repo.GetTrainingByID(ctx, training.ID)
	if err != nil {
		return models.TrainingPlan{}, err
	}

	training.Version = oldPlan.Version + 1
	training.MeanScore = oldPlan.MeanScore
	training.FavoritesCount = oldPlan.FavoritesCount
	training.SessionsCount = oldPlan.SessionsCount
	training.Reviews = oldPlan.Reviews

	deleteOldResult := db.Select("Exercises").Delete(&models.TrainingPlan{ID: oldPlan.ID, Version: oldPlan.Version})
	if deleteOldResult.Error != nil {
		repo.logger.Error("Unable to delete training plan", zap.Error(deleteOldResult.Error))
		return models.TrainingPlan{}, deleteOldResult.Error
	}

	result := db.Create(&training)
	if result.Error != nil {
		repo.logger.Error("Unable to update training plan", zap.Error(result.Error), zap.Any("training", training))
		return models.TrainingPlan{}, result.Error
	}
	return training, nil
}

func (repo TrainingRepository) UpdateDisabledStatus(ctx context.Context, trainingID uint, disabled bool) error {
	db := repo.db.WithContext(ctx)

	training := models.TrainingPlan{}
	if err := db.First(&training, trainingID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return contracts.ErrTrainingPlanNotFound
		}
		repo.logger.Error("Unable to find training plan", zap.Error(err))
		return err
	}

	if training.Disabled == disabled {
		if disabled {
			return contracts.ErrTrainingAlreadyDisabled
		}
		return contracts.ErrTrainingNotDisabled
	}

	result := db.Model(&training).Update("disabled", disabled)
	if result.Error != nil {
		repo.logger.Error("Unable to update training plan", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func (repo TrainingRepository) DeleteTrainingPlan(ctx context.Context, trainingID uint) error {
	db := repo.db.WithContext(ctx)
	result := db.Select("Exercises", "Reviews", "Favorites").Delete(&models.TrainingPlan{ID: trainingID})
	if result.Error != nil {
		repo.logger.Error("Unable to delete training plan", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrTrainingPlanNotFound
	}
	return nil
}

func (repo TrainingRepository) AddToFavorite(ctx context.Context, userID string, trainingID uint, trainingVersion uint) error {
	db := repo.db.WithContext(ctx)
	result := db.Create(&models.Favorite{UserID: userID, TrainingPlanID: trainingID, TrainingPlanVersion: trainingVersion})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return contracts.ErrAlreadyLiked
		}
		repo.logger.Error("Unable to add training plan to favorites", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func (repo TrainingRepository) RemoveFromFavorite(ctx context.Context, userID string, trainingID uint, trainingVersion uint) error {
	db := repo.db.WithContext(ctx)
	result := db.Where("user_id = ? AND training_plan_id = ?", userID, trainingID).Delete(&models.Favorite{})
	if result.Error != nil {
		repo.logger.Error("Unable to remove training plan from favorites", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected < 1 {
		return contracts.ErrNotLiked
	}
	return nil
}

func (repo TrainingRepository) GetFavoriteTrainings(ctx context.Context, req trainings.GetFavoritesRequest) (trainings.GetTrainingsResponse, error) {
	var res []models.TrainingPlan
	db := repo.db.WithContext(ctx)
	result := db.Where("training_plans.id IN (SELECT training_plan_id FROM favorites WHERE user_id = ?) AND training_plans.disabled = false", req.UserID).
		Scopes(database.Paginate(&res, &req.Pagination, db)).
		Preload("Exercises").
		Preload("Reviews").
		Preload("Tags").
		Select(`training_plans.*, COALESCE((SELECT AVG(score) FROM reviews WHERE reviews.training_plan_id = training_plans.id), 0) as mean_score, 
					(SELECT COUNT(*) FROM favorites WHERE favorites.training_plan_id = training_plans.id) AS favorites_count, 
					(SELECT COUNT(*) FROM training_sessions WHERE training_sessions.training_plan_id = training_plans.id) as sessions_count`).
		Find(&res)
	if result.Error != nil {
		repo.logger.Error("Unable to get favorite trainings", zap.Error(result.Error))
		return trainings.GetTrainingsResponse{}, result.Error
	}
	var results []Result
	result.Scan(&results)

	for i := range res {
		res[i].MeanScore = results[i].MeanScore
		res[i].FavoritesCount = results[i].FavoritesCount
		res[i].SessionsCount = results[i].SessionsCount
	}
	return trainings.GetTrainingsResponse{TrainingPlans: res, Pagination: req.Pagination}, nil
}

type Result struct {
	models.TrainingPlan
	MeanScore      float32
	FavoritesCount uint
	SessionsCount  uint
}
