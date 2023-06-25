package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/users"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddFavorite struct {
	trainings utrainings.FavoriteAdder
	logger    *zap.Logger
}

func NewAddFavorite(trainings utrainings.FavoriteAdder, logger *zap.Logger) AddFavorite {
	return AddFavorite{trainings: trainings, logger: logger}
}

func (h AddFavorite) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req users.UserID
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		trainingID := ctx.MustGet("trainingID").(uint)

		err = h.trainings.AddToFavorite(ctx, req.UserID, trainingID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
