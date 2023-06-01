package repositories

import (
	"context"
	"strings"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TrainingSessions interface {
	Create(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error)
	Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) ([]models.TrainingSession, error)
	Update(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error)
}

type TrainingSessionsRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (repo TrainingSessionsRepository) Create(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error) {
	db := repo.db.WithContext(ctx)
	res := db.Create(&session)

	if res.Error != nil {
		if strings.Contains(res.Error.Error(), contracts.ErrForeignKey.Error()) {
			return models.TrainingSession{}, contracts.ErrTrainingPlanNotFound
		}
		repo.logger.Error(res.Error.Error(), zap.Any("session", session))
		return models.TrainingSession{}, res.Error
	}

	return session, nil
}

func (repo TrainingSessionsRepository) Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) ([]models.TrainingSession, error) {
	db := repo.db.WithContext(ctx)

	var sessions []models.TrainingSession

	db = db.Where("user_id = ?", req.UserID)

	if req.TrainingID != 0 {
		db = db.Where("training_id = ?", req.TrainingID)
	}

	res := db.Order("updated_at desc").Preload("TrainingPlans").Preload("ExerciseSessions").Preload("Exercises").Find(&sessions)

	if res.Error != nil {
		repo.logger.Error(res.Error.Error(), zap.Any("req", req))
		return nil, res.Error
	}
	return sessions, nil
}
