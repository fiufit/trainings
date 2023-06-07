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

type CreateGoal struct {
	goals  ugoals.GoalCreator
	logger *zap.Logger
}

func NewCreateGoal(goals ugoals.GoalCreator, logger *zap.Logger) CreateGoal {
	return CreateGoal{goals: goals, logger: logger}
}

func (h CreateGoal) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req goals.CreateGoalRequest
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
		err = goals.ValidateGoalType(req.GoalType, req.GoalSubtype)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(err))
			return
		}
		res, err := h.goals.CreateGoal(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrUserNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
