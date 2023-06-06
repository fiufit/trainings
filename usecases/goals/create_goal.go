package goals

import (
	"context"

	"github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type GoalCreator interface {
	CreateGoal(ctx context.Context, req goals.CreateGoalRequest) (models.Goal, error)
}

type GoalCreatorImpl struct {
	users  repositories.Users
	goals  repositories.Goals
	logger *zap.Logger
}

func NewGoalCreatorImpl(users repositories.Users, goals repositories.Goals, logger *zap.Logger) GoalCreatorImpl {
	return GoalCreatorImpl{users: users, goals: goals, logger: logger}
}

func (uc *GoalCreatorImpl) CreateGoal(ctx context.Context, req goals.CreateGoalRequest) (models.Goal, error) {
	_, err := uc.users.GetUserByID(ctx, req.UserID)
	if err != nil {
		return models.Goal{}, err
	}
	newGoal := models.Goal{
		Title:             req.Title,
		GoalValue:         req.GoalValue,
		GoalValueProgress: 0,
		GoalType:          req.GoalType,
		GoalSubtype:       req.GoalSubtype,
		Deadline:          req.Deadline,
		UserID:            req.UserID,
	}
	createdGoal, err := uc.goals.Create(ctx, newGoal)
	return createdGoal, err
}
