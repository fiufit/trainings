package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	ureviews "github.com/fiufit/trainings/usecases/reviews"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetReviews struct {
	reviews ureviews.ReviewGetter
	logger  *zap.Logger
}

func NewGetReviews(reviews ureviews.ReviewGetter, logger *zap.Logger) GetReviews {
	return GetReviews{reviews: reviews, logger: logger}
}

func (h GetReviews) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reviews.GetReviewsRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		req.Pagination.Validate()
		trainingID := ctx.MustGet("trainingID").(uint)
		req.TrainingPlanID = trainingID
		resReviews, err := h.reviews.GetReviews(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resReviews))
	}
}
