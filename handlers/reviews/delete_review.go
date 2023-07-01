package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	ureviews "github.com/fiufit/trainings/usecases/reviews"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteReview struct {
	reviews ureviews.ReviewDeleter
	logger  *zap.Logger
}

func NewDeleteReview(reviews ureviews.ReviewDeleter, logger *zap.Logger) DeleteReview {
	return DeleteReview{reviews: reviews, logger: logger}
}

// Delete review godoc
//	@Summary		Deletes a review
//	@Description	Deletes a review for a given reviewID and trainingID given by the userID
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training Plan ID"
// 	@Param			reviewID				path		uint								true	"Review ID"
//	@Param			payload					body		reviews.DeleteReviewRequest	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
// 	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/reviews/{reviewID} 	[delete]
func (h DeleteReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reviews.DeleteReviewRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		trainingID := ctx.MustGet("trainingID").(uint)
		reviewID := ctx.MustGet("reviewID").(uint)
		req.TrainingPlanID = trainingID
		req.ReviewID = reviewID

		err = h.reviews.DeleteReview(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
