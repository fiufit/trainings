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

// Get goals godoc
//	@Summary		Get goals by different query params with pagination
//	@Description	Get goals  with pagination filtered by their deadline, type, subtype for a given userID
//	@Tags			goals
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
//	@Param			userID					query		string 							true	"User ID"
//	@Param			type					query		string 							false	"Goal Type"
//	@Param			subtype					query		string 							false	"Goal Subtype"
//	@Param			deadline				query		string 							false	"Goal Deadline"
//	@Success		200						{object}	goals.GetGoalsResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/goals 	[get]
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
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resGoals))
	}
}
