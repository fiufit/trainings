package handlers

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateTraining struct {
	trainings usecases.TrainingUpdater
	logger    *zap.Logger
}

func NewUpdateTraining(trainings usecases.TrainingUpdater, logger *zap.Logger) UpdateTraining {
	return UpdateTraining{trainings: trainings, logger: logger}
}

func (h UpdateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req training.UpdateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(string)
		req.ID = trainingID

		updatedTraining, err := h.trainings.UpdateTrainingPlan(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrTrainingPlanNotFound))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedTraining))
	}
}
