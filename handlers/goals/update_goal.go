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

type UpdateGoal struct {
	goals  ugoals.GoalUpdater
	logger *zap.Logger
}

func NewUpdateGoal(goals ugoals.GoalUpdater, logger *zap.Logger) UpdateGoal {
	return UpdateGoal{goals: goals, logger: logger}
}

func (h UpdateGoal) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req goals.UpdateGoalRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(err))
			return
		}
		goalID := ctx.MustGet("goalID").(uint)
		req.GoalID = goalID

		updatedGoal, err := h.goals.UpdateGoal(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrGoalNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrGoalNotFound))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedAthlete) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(contracts.ErrUnauthorizedAthlete))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedGoal))
	}
}
