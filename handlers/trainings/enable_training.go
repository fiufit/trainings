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
