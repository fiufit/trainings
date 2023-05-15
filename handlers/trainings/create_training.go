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

type CreateTraining struct {
	trainings utrainings.TrainingCreator
	logger    *zap.Logger
}

func NewCreateTraining(trainings utrainings.TrainingCreator, logger *zap.Logger) CreateTraining {
	return CreateTraining{trainings: trainings, logger: logger}
}

func (h CreateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.CreateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		res, err := h.trainings.CreateTraining(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrUserNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
