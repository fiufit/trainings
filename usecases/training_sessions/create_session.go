package training_sessions

import (
	"context"
	"time"

	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type TrainingSessionCreator interface {
	CreateTrainingSession(ctx context.Context, trainingID uint, userID string) (tsContracts.CreateTrainingSessionResponse, error)
}

type TrainingSessionCreatorImpl struct {
	users     repositories.Users
	trainings repositories.TrainingPlans
	sessions  repositories.TrainingSessions
	logger    *zap.Logger
}

func (uc *TrainingSessionCreatorImpl) CreateTrainingSession(ctx context.Context, trainingID uint, userID string) (tsContracts.CreateTrainingSessionResponse, error) {
	_, err := uc.users.GetUserByID(ctx, userID)
	if err != nil {
		return tsContracts.CreateTrainingSessionResponse{}, err
	}

	training, err := uc.trainings.GetTrainingByID(ctx, trainingID)
	if err != nil {
		return tsContracts.CreateTrainingSessionResponse{}, err
	}

	session := makeTrainingSession(training, userID)
	return uc.sessions.Create(ctx, session)
}

func makeTrainingSession(training models.TrainingPlan, userID string) models.TrainingSession {
	exerciseSessions := make([]models.ExerciseSession, len(training.Exercises))
	for i, exercise := range training.Exercises {
		exerciseSessions[i] = models.ExerciseSession{
			ExerciseID: exercise.ID,
			Exercise:   exercise,
			Done:       false,
		}
	}

	trainingSession := models.TrainingSession{
		TrainingPlanID:      training.ID,
		TrainingPlanVersion: training.Version,
		TrainingPlan:        training,
		UserID:              userID,
		ExerciseSessions:    exerciseSessions,
		Done:                false,
		StepCount:           0,
		SecondsCount:        0,
		UpdatedAt:           time.Now(),
	}

	return trainingSession
}
