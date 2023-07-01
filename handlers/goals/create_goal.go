package handlers

import (
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

// Create goal godoc
//	@Summary		Creates a new goal
//	@Description	Creates a new training goal for a given userID
//	@Tags			goals
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
//	@Param			payload					body		goals.CreateGoalRequest	true	"Body params"
//	@Success		200						{object}	models.Goal	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/goals 	[post]
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
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
