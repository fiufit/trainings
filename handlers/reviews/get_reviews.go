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

// Get reviews godoc
//	@Summary		Get reviews by different query params with pagination
//	@Description	Get reviews of a training plan with pagination filtered by their score, comment or userID
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training Plan ID"
// 	@Param			min_score				query		uint		false	"minimum score of the review"
// 	@Param			max_score				query		uint		false	"maximum score of the review"
// 	@Param			comment					query		string	false	"comment of the review"
// 	@Param			user_id					query		string	false	"userID of the review"
//	@Param			page				query		int		false	"page number when getting with pagination"
//	@Param			page_size			query		int		false	"page size when getting with pagination"
//	@Success		200						{object}	reviews.GetReviewsResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/reviews	[get]
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
