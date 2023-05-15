package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetTrainings struct {
	trainings utrainings.TrainingGetter
	logger    *zap.Logger
}

func NewGetTrainings(trainings utrainings.TrainingGetter, logger *zap.Logger) GetTrainings {
	return GetTrainings{trainings: trainings, logger: logger}
}

func (h GetTrainings) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.GetTrainingsRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		req.Pagination.Validate()
		resTrainings, err := h.trainings.GetTrainingPlans(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resTrainings))
	}
}