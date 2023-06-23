package training_sessions

import (
	"context"
	"time"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/metrics"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingSessionUpdater interface {
	UpdateTrainingSession(ctx context.Context, req tsContracts.UpdateTrainingSessionRequest) (tsContracts.UpdateTrainingSessionResponse, error)
}

type TrainingSessionUpdaterImpl struct {
	sessions      repositories.TrainingSessions
	firebase      repositories.Firebase
	goals         repositories.Goals
	notifications repositories.Notifications
	metrics       repositories.Metrics
	logger        *zap.Logger
}

func NewTrainingSessionUpdaterImpl(sessions repositories.TrainingSessions, firebase repositories.Firebase, goals repositories.Goals, notifications repositories.Notifications, metrics repositories.Metrics, logger *zap.Logger) TrainingSessionUpdaterImpl {
	return TrainingSessionUpdaterImpl{sessions: sessions, firebase: firebase, goals: goals, notifications: notifications, metrics: metrics, logger: logger}
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

	if updatedSession.Done {

		sessionDoneMetric := metrics.CreateMetricRequest{
			MetricType: "training_session_finished",
			SubType:    ts.UserID,
		}

		uc.metrics.Create(ctx, sessionDoneMetric)

		goals, err := uc.goals.UpdateBySession(ctx, updatedSession)
		if err != nil {
			uc.logger.Error("Unable to update user goal", zap.Error(err))
		} else {
			for _, goal := range goals {
				if goal.GoalValue <= goal.GoalValueProgress {
					err = uc.notifications.SendGoalNotification(ctx, goal)
					if err != nil {
						uc.logger.Error("Unable to send notification", zap.Error(err))
					}
				}

			}
		}
	}

	uc.firebase.FillTrainingPicture(ctx, &updatedSession.TrainingPlan)

	return tsContracts.UpdateTrainingSessionResponse{Session: updatedSession}, nil
}
