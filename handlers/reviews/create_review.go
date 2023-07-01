package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/reviews"
	ureviews "github.com/fiufit/trainings/usecases/reviews"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateReview struct {
	reviews ureviews.ReviewCreator
	logger  *zap.Logger
}

func NewCreateReview(reviews ureviews.ReviewCreator, logger *zap.Logger) CreateReview {
	return CreateReview{reviews: reviews, logger: logger}
}

// Create review godoc
//	@Summary		Creates a review for a training plan
//	@Description	Creates a new review with a comment and a score for a given trainingPlanID and userID
//	@Tags			reviews
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training Plan ID"
//	@Param			payload					body		reviews.CreateReviewRequest	true	"Body params"
//	@Success		200						{object}	models.Review	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
// 	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
// 	@Failure		409						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/reviews 	[post]
func (h CreateReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req reviews.CreateReviewRequest
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
		req.TrainingPlanID = trainingID

		res, err := h.reviews.CreateReview(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
