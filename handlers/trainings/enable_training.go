package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EnableTraining struct {
	trainings utrainings.TrainingUpdater
	logger    *zap.Logger
}

func NewEnableTraining(trainings utrainings.TrainingUpdater, logger *zap.Logger) EnableTraining {
	return EnableTraining{trainings: trainings, logger: logger}
}

// Enable training plan godoc
//	@Summary		Enables a training plan
//	@Description	Enables a training plan that had been disabled previously given its ID
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training plan ID"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
// 	@Failure		409						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/enable	[post]
func (h EnableTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trainingID := ctx.MustGet("trainingID").(uint)
		err := h.trainings.EnableTrainingPlan(ctx, trainingID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
