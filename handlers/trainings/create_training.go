package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateTraining struct {
	trainings utrainings.TrainingCreator
	logger    *zap.Logger
}

func NewCreateTraining(trainings utrainings.TrainingCreator, logger *zap.Logger) CreateTraining {
	return CreateTraining{trainings: trainings, logger: logger}
}

// Create training plan godoc
//	@Summary		Creates a new training plan
//	@Description	Creates a new training plan with a name, description, difficulty, duration, exercises and tags
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
//	@Param			payload					body		trainings.CreateTrainingRequest	true	"Body params"
//	@Success		200						{object}	trainings.CreateTrainingResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings	[post]
func (h CreateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.CreateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		validateErr := req.Validate()
		if err != nil || validateErr != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		res, err := h.trainings.CreateTraining(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
