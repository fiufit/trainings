package handlers

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/goals"
	ugoals "github.com/fiufit/trainings/usecases/goals"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteGoal struct {
	goals  ugoals.GoalDeleter
	logger *zap.Logger
}

func NewDeleteGoal(goals ugoals.GoalDeleter, logger *zap.Logger) DeleteGoal {
	return DeleteGoal{goals: goals, logger: logger}
}

func (h DeleteGoal) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req goals.DeleteGoalRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		goalID := ctx.MustGet("goalID").(uint)
		req.GoalID = goalID

		err = h.goals.DeleteGoal(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrGoalNotFound))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedAthlete) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrUnauthorizedAthlete))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
