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
