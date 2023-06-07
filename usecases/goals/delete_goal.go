package goals

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type GoalDeleter interface {
	DeleteGoal(ctx context.Context, req goals.DeleteGoalRequest) error
}

type GoalDeleterImpl struct {
	goals  repositories.Goals
	logger *zap.Logger
}

func NewGoalDeleterImpl(goals repositories.Goals, logger *zap.Logger) GoalDeleterImpl {
	return GoalDeleterImpl{goals: goals, logger: logger}
}

func (uc *GoalDeleterImpl) DeleteGoal(ctx context.Context, req goals.DeleteGoalRequest) error {
	goal, err := uc.goals.GetByID(ctx, req.GoalID)
	if err != nil {
		return err
	}
	if goal.UserID != req.UserID {
		return contracts.ErrUnauthorizedAthlete
	}
	return uc.goals.Delete(ctx, req.GoalID)
}
