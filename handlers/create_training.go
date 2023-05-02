package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateTraining struct {
	trainings usecases.TrainingCreator
	logger    *zap.Logger
}

func NewCreateTraining(trainings usecases.TrainingCreator, logger *zap.Logger) CreateTraining {
	return CreateTraining{trainings: trainings, logger: logger}
}

func (h CreateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req training.CreateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		res, err := h.trainings.CreateTraining(ctx, req)
		if err != nil {
			h.logger.Error("unable to create new training", zap.Error(err), zap.Any("request", req))
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
