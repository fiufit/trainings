package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetTrainings struct {
	trainings utrainings.TrainingGetter
	logger    *zap.Logger
}

func NewGetTrainings(trainings utrainings.TrainingGetter, logger *zap.Logger) GetTrainings {
	return GetTrainings{trainings: trainings, logger: logger}
}

// Get training plans godoc
//	@Summary		Get traning plans with pagination
//	@Description	Get training plans with pagination, optionally filtered by name, description, difficulty, duration and tags
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
//	@Param			name					query		string							false	"Name"
//	@Param			description				query		string							false	"Description"
//	@Param			difficulty				query		string							false	"Difficulty"
//	@Param			trainer_id				query		string							false	"Trainer ID"
//	@Param			min_duration			query		uint							false	"Min Duration"
//	@Param			max_duration			query		uint							false	"Max Duration"
//	@Param			tags					query		[]string						false	"Tags"
//	@Param			disabled				query		bool							false	"Disabled"
//	@Param			page					query		uint							false	"Page"
//	@Param			page_size				query		uint							false	"Page Size"
//	@Param			user_id					query		string							false	"User ID"
//	@Success		200						{object}	trainings.GetTrainingsResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings	[get]
func (h GetTrainings) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.GetTrainingsRequest
		err := ctx.ShouldBindQuery(&req)
		validateErr := req.Validate()
		if err != nil || validateErr != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		req.Pagination.Validate()
		var resTrainings trainings.GetTrainingsResponse
		if req.UserID != "" {
			resTrainings, err = h.trainings.GetRecommendedPlans(ctx, req)
		} else {
			resTrainings, err = h.trainings.GetTrainingPlans(ctx, req)
		}

		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(resTrainings))
	}
}
