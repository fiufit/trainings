package training_sessions

import (
	"context"
	"time"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingSessionUpdater interface {
	UpdateTrainingSession(ctx context.Context, req tsContracts.UpdateTrainingSessionRequest) (tsContracts.UpdateTrainingSessionResponse, error)
}

type TrainingSessionUpdaterImpl struct {
	sessions repositories.TrainingSessions
	firebase repositories.Firebase
	logger   *zap.Logger
}

func NewTrainingSessionUpdaterImpl(sessions repositories.TrainingSessions, firebase repositories.Firebase, logger *zap.Logger) TrainingSessionUpdaterImpl {
	return TrainingSessionUpdaterImpl{sessions: sessions, firebase: firebase, logger: logger}
}

func (uc *TrainingSessionUpdaterImpl) UpdateTrainingSession(ctx context.Context, req tsContracts.UpdateTrainingSessionRequest) (tsContracts.UpdateTrainingSessionResponse, error) {
	ts, err := uc.sessions.GetByID(ctx, req.ID)
	if err == nil && ts.UserID != req.RequesterID {
		return tsContracts.UpdateTrainingSessionResponse{}, contracts.ErrUnauthorizedAthlete
	}
	if err != nil {
		return tsContracts.UpdateTrainingSessionResponse{}, err
	}

	if ts.Done {
		return tsContracts.UpdateTrainingSessionResponse{}, contracts.ErrTrainingSessionAlreadyFinished
	}

	for _, exerciseSessionRequest := range req.ExerciseSessions {

		for i, exerciseSession := range ts.ExerciseSessions {
			if exerciseSession.ID == exerciseSessionRequest.ID {

				if *req.Done && !*exerciseSessionRequest.Done {
					return tsContracts.UpdateTrainingSessionResponse{}, contracts.ErrTrainingSessionNotComplete
				}

				ts.ExerciseSessions[i].Done = *exerciseSessionRequest.Done
			}
		}
	}

	ts.StepCount = *req.StepCount
	ts.SecondsCount = *req.SecondsCount
	ts.Done = *req.Done
	ts.UpdatedAt = time.Now()

	updatedSession, err := uc.sessions.Update(ctx, ts)
	if err != nil {
		return tsContracts.UpdateTrainingSessionResponse{}, err
	}

	uc.firebase.FillTrainingPicture(ctx, &updatedSession.TrainingPlan)

	return tsContracts.UpdateTrainingSessionResponse{Session: updatedSession}, nil
}
