package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetFavorites struct {
	trainings utrainings.TrainingGetter
	logger    *zap.Logger
}

func NewGetFavorites(trainings utrainings.TrainingGetter, logger *zap.Logger) GetFavorites {
	return GetFavorites{trainings: trainings, logger: logger}
}

func (h GetFavorites) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.GetFavoritesRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		req.Pagination.Validate()
		resTrainings, err := h.trainings.GetFavoritePlans(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resTrainings))
	}
}
