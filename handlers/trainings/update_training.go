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

type UpdateTraining struct {
	trainings utrainings.TrainingUpdater
	logger    *zap.Logger
}

func NewUpdateTraining(trainings utrainings.TrainingUpdater, logger *zap.Logger) UpdateTraining {
	return UpdateTraining{trainings: trainings, logger: logger}
}

func (h UpdateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.UpdateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		validateErr := req.Validate()
		if err != nil || validateErr != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		req.ID = trainingID

		updatedTraining, err := h.trainings.UpdateTrainingPlan(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrTrainingPlanNotFound))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedTrainer) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(contracts.ErrUnauthorizedTrainer))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedTraining))
	}
}
