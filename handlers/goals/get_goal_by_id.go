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

// Get goal godoc
//	@Summary		Gets a goal by ID
//	@Description	Gets a goal for a given goalID
//	@Tags			goals
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			goalID					path		uint								true	"Goal ID"
//	@Success		200						{object}	models.Goal	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/goals/{goalID} 	[get]
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
