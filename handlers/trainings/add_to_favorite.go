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

// Add to favorites godoc
//	@Summary		Adds a training plan to favorites
//	@Description	Adds a training plan to favorites for a user given its ID
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training plan ID"
//	@Param			payload					body		users.UserID	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
// 	@Failure		409						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/favorites	[post]
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
