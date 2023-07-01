package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DisableTraining struct {
	trainings utrainings.TrainingUpdater
	logger    *zap.Logger
}

func NewDisableTraining(trainings utrainings.TrainingUpdater, logger *zap.Logger) DisableTraining {
	return DisableTraining{trainings: trainings, logger: logger}
}

// Disable training plan godoc
//	@Summary		Disables a training plan
//	@Description	Disables a training plan that had been enabled previously (or had never been disabled) given its ID
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
//	@Router			/{version}/trainings/{trainingID}/disable	[delete]
func (h DisableTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trainingID := ctx.MustGet("trainingID").(uint)
		err := h.trainings.DisableTrainingPlan(ctx, trainingID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
