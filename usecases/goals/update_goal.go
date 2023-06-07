package goals

import (
	"context"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type GoalUpdater interface {
	UpdateGoal(ctx context.Context, req goals.UpdateGoalRequest) (models.Goal, error)
}

type GoalUpdaterImpl struct {
	goals  repositories.Goals
	logger *zap.Logger
}

func NewGoalUpdaterImpl(goals repositories.Goals, logger *zap.Logger) GoalUpdaterImpl {
	return GoalUpdaterImpl{goals: goals, logger: logger}
}

func (uc *GoalUpdaterImpl) UpdateGoal(ctx context.Context, req goals.UpdateGoalRequest) (models.Goal, error) {
	goal, err := uc.goals.GetByID(ctx, req.GoalID)
	if err != nil {
		return models.Goal{}, err
	}

	if goal.UserID != req.UserID {
		return models.Goal{}, contracts.ErrUnauthorizedAthlete
	}

	patchedGoal := uc.patchGoalModel(ctx, goal, req)

	updatedGoal, err := uc.goals.Update(ctx, patchedGoal)
	if err != nil {
		return models.Goal{}, err
	}

	return updatedGoal, nil
}

func (uc *GoalUpdaterImpl) patchGoalModel(ctx context.Context, goal models.Goal, req goals.UpdateGoalRequest) models.Goal {
	if req.Title != "" {
		goal.Title = req.Title
	}

	if req.GoalValue != 0 {
		goal.GoalValue = req.GoalValue
	}

	if !req.Deadline.IsZero() {
		goal.Deadline = req.Deadline
	}

	return goal
}
