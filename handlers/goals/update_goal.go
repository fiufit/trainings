package handlers

import (
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

// Update goal godoc
//	@Summary		Updates a goal
//	@Description	Updates a goal for a given goalID, validating that the goal was created by the user
//	@Tags			goals
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			goalID					path		uint								true	"Goal ID"
//	@Param			payload					body		goals.UpdateGoalRequest	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
// 	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/goals/{goalID} 	[patch]
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
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedGoal))
	}
}
