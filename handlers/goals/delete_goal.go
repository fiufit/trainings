package handlers

import (
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

// Delete goal godoc
//	@Summary		Deletes a goal
//	@Description	Delete a goal for a given goalID, validating that the goal belongs to the userID
//	@Tags			goals
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			goalID					path		uint								true	"Goal ID"
//	@Param			payload					body		goals.DeleteGoalRequest	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
// 	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/goals/{goalID} 	[delete]
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
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
