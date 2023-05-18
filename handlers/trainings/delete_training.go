package handlers

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteTraining struct {
	trainings utrainings.TrainingDeleter
	logger    *zap.Logger
}

func NewDeleteTraining(trainings utrainings.TrainingDeleter, logger *zap.Logger) DeleteTraining {
	return DeleteTraining{trainings: trainings, logger: logger}
}

func (h DeleteTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.DeleteTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		req.TrainingPlanID = trainingID

		err = h.trainings.DeleteTraining(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrUnauthorizedTrainer) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(err))
				return
			}
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
