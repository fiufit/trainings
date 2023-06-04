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
	firebase repositories.Firebase
	sessions repositories.TrainingSessions
	logger   *zap.Logger
}

func NewTrainingSessionGetterImpl(sessions repositories.TrainingSessions, firebase repositories.Firebase, logger *zap.Logger) TrainingSessionGetterImpl {
	return TrainingSessionGetterImpl{sessions: sessions, firebase: firebase, logger: logger}
}

func (uc *TrainingSessionGetterImpl) GetByID(ctx context.Context, sessionID uint, requesterID string) (models.TrainingSession, error) {

	session, err := uc.sessions.GetByID(ctx, sessionID)

	if err == nil && session.UserID != requesterID {
		return models.TrainingSession{}, contracts.ErrUnauthorizedAthlete
	}

	if err != nil {
		return models.TrainingSession{}, err
	}

	uc.firebase.FillTrainingPicture(ctx, &session.TrainingPlan)

	return session, nil
}

func (uc *TrainingSessionGetterImpl) Get(ctx context.Context, req tsContracts.GetTrainingSessionsRequest) (tsContracts.GetTrainingSessionsResponse, error) {
	sessions, err := uc.sessions.Get(ctx, req)
	if err != nil {
		return tsContracts.GetTrainingSessionsResponse{}, err
	}

	for i, _ := range sessions {
		uc.firebase.FillTrainingPicture(ctx, &(sessions[i].TrainingPlan))
	}

	res := tsContracts.GetTrainingSessionsResponse{
		Sessions:   sessions,
		Pagination: req.Pagination,
	}

	return res, nil
}
