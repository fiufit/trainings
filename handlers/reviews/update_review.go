package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	ureviews "github.com/fiufit/trainings/usecases/reviews"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateReview struct {
	reviews ureviews.ReviewUpdater
	logger  *zap.Logger
}

func NewUpdateReview(reviews ureviews.ReviewUpdater, logger *zap.Logger) UpdateReview {
	return UpdateReview{reviews: reviews, logger: logger}
}

func (h UpdateReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reviews.UpdateReviewRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(err))
			return
		}
		trainingID := ctx.MustGet("trainingID").(uint)
		reviewID := ctx.MustGet("reviewID").(uint)
		req.TrainingPlanID = trainingID
		req.ReviewID = reviewID

		updatedReview, err := h.reviews.UpdateReview(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedReview))
	}
}
