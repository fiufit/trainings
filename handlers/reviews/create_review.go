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
