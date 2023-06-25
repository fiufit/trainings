package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	ugoals "github.com/fiufit/trainings/usecases/goals"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetGoalByID struct {
	goals  ugoals.GoalGetter
	logger *zap.Logger
}

func NewGetGoalByID(goals ugoals.GoalGetter, logger *zap.Logger) GetGoalByID {
	return GetGoalByID{goals: goals, logger: logger}
}

func (h GetGoalByID) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		goalID := ctx.MustGet("goalID").(uint)
		goal, err := h.goals.GetGoalByID(ctx, goalID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(goal))
	}
}
