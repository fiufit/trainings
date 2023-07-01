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

// Get review godoc
//	@Summary		Gets a review by ID
//	@Description	Gets a review for a given reviewID
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training Plan ID"
// 	@Param			reviewID				path		uint								true	"Review ID"
//	@Success		200						{object}	models.Review	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/reviews/{reviewID} 	[get]
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
