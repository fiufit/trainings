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

// Get favorites godoc
//	@Summary		Gets favorite training plans for a user
//	@Description	Gets favorite training plans for a user with pagination
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			user_id					query		string								true	"User ID"
//	@Param			page					query		uint							false	"Page"
//	@Param			page_size				query		uint							false	"Page Size"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/favorites	[get]
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
