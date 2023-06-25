package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	ureviews "github.com/fiufit/trainings/usecases/reviews"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetReviewByID struct {
	reviews ureviews.ReviewGetter
	logger  *zap.Logger
}

func NewGetReviewByID(reviews ureviews.ReviewGetter, logger *zap.Logger) GetReviewByID {
	return GetReviewByID{reviews: reviews, logger: logger}
}

func (h GetReviewByID) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trainingID := ctx.MustGet("trainingID").(uint)
		reviewID := ctx.MustGet("reviewID").(uint)
		review, err := h.reviews.GetReviewByID(ctx, trainingID, reviewID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(review))
	}
}
