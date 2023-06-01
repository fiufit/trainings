package repositories

import (
	"context"
	"errors"
	"strings"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TrainingSessions interface {
	Create(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error)
	Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) ([]models.TrainingSession, error)
	GetByID(ctx context.Context, sessionID uint) (models.TrainingSession, error)
	Update(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error)
}

type TrainingSessionsRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTrainingSessionsRepository(db *gorm.DB, logger *zap.Logger) TrainingSessionsRepository {
	return TrainingSessionsRepository{db: db, logger: logger}
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

func (repo TrainingSessionsRepository) GetByID(ctx context.Context, sessionID uint) (models.TrainingSession, error) {
	db := repo.db.WithContext(ctx)
	var session models.TrainingSession
	res := db.Preload("TrainingPlan").
		Preload("TrainingPlan.Exercises").
		Preload("ExerciseSessions").
		Preload("ExerciseSessions.Exercise").
		Preload("TrainingPlan.Tags").
		First(&session, "id = ?", sessionID)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return models.TrainingSession{}, contracts.ErrTrainingSessionNotFound
		}

		repo.logger.Error(res.Error.Error(), zap.Any("sessionID", sessionID))
		return models.TrainingSession{}, res.Error
	}

	return session, nil
}

func (repo TrainingSessionsRepository) Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) ([]models.TrainingSession, error) {
	var sessions []models.TrainingSession
	db := repo.db.WithContext(ctx)

	db = db.Where("user_id = ?", req.UserID)

	if req.TrainingID != 0 {
		db = db.Where("training_plan_id = ?", req.TrainingID)
	}

	res := db.Unscoped().Scopes(database.Paginate(sessions, &req.Pagination, db)).Order("updated_at desc").
		Preload("TrainingPlan").
		Preload("TrainingPlan.Exercises").
		Preload("ExerciseSessions").
		Preload("ExerciseSessions.Exercise").
		Preload("TrainingPlan.Tags").
		Find(&sessions)
	if res.Error != nil {
		repo.logger.Error(res.Error.Error(), zap.Any("req", req))
		return nil, res.Error
	}
	return sessions, nil
}

func (repo TrainingSessionsRepository) Update(ctx context.Context, session models.TrainingSession) (models.TrainingSession, error) {
	db := repo.db.WithContext(ctx)
	res := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&session)

	if res.Error != nil {
		repo.logger.Error(res.Error.Error(), zap.Any("session", session))
		return models.TrainingSession{}, res.Error
	}

	return session, nil
}
