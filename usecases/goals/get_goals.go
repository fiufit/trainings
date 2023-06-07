package goals

import (
	"context"

	"github.com/fiufit/trainings/contracts/goals"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"go.uber.org/zap"
)

type GoalGetter interface {
	GetGoals(ctx context.Context, req goals.GetGoalsRequest) (goals.GetGoalsResponse, error)
	GetGoalByID(ctx context.Context, goalID uint) (models.Goal, error)
}

type GoalGetterImpl struct {
	goals  repositories.Goals
	logger *zap.Logger
}

func NewGoalGetterImpl(goals repositories.Goals, logger *zap.Logger) GoalGetterImpl {
	return GoalGetterImpl{goals: goals, logger: logger}
}

func (uc *GoalGetterImpl) GetGoalByID(ctx context.Context, goalID uint) (models.Goal, error) {
	return uc.goals.GetByID(ctx, goalID)
}

func (uc *GoalGetterImpl) GetGoals(ctx context.Context, req goals.GetGoalsRequest) (goals.GetGoalsResponse, error) {
	res, err := uc.goals.GetByUserID(ctx, req.UserID)
	return goals.GetGoalsResponse{Goals: res}, err
}
