package training_sessions

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingSessionGetter interface {
	GetByID(ctx context.Context, sessionID uint, requesterID string) (models.TrainingSession, error)
	Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) (tsContracts.GetTrainingSessionsResponse, error)
}

type TrainingSessionGetterImpl struct {
	sessions repositories.TrainingSessions
	logger   *zap.Logger
}

func NewTrainingSessionGetterImpl(sessions repositories.TrainingSessions, logger *zap.Logger) TrainingSessionGetterImpl {
	return TrainingSessionGetterImpl{sessions: sessions, logger: logger}
}

func (uc *TrainingSessionGetterImpl) GetByID(ctx context.Context, sessionID uint, requesterID string) (models.TrainingSession, error) {

	session, err := uc.sessions.GetByID(ctx, sessionID)

	if err == nil && session.UserID != requesterID {
		return models.TrainingSession{}, contracts.ErrUnauthorizedAthlete
	}

	if err != nil {
		return models.TrainingSession{}, err
	}
	return session, nil
}

func (uc *TrainingSessionGetterImpl) Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) (tsContracts.GetTrainingSessionsResponse, error) {
	sessions, err := uc.sessions.Get(ctx, req)
	if err != nil {
		return tsContracts.GetTrainingSessionsResponse{}, err
	}

	res := tsContracts.GetTrainingSessionsResponse{
		Sessions:   sessions,
		Pagination: req.Pagination,
	}

	return res, nil
}
