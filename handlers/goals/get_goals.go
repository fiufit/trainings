package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/goals"
	ugoals "github.com/fiufit/trainings/usecases/goals"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetGoals struct {
	goals  ugoals.GoalGetter
	logger *zap.Logger
}

func NewGetGoals(goals ugoals.GoalGetter, logger *zap.Logger) GetGoals {
	return GetGoals{goals: goals, logger: logger}
}

func (h GetGoals) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req goals.GetGoalsRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		resGoals, err := h.goals.GetGoals(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resGoals))
	}
}
