package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateTraining struct {
	trainings utrainings.TrainingUpdater
	logger    *zap.Logger
}

func NewUpdateTraining(trainings utrainings.TrainingUpdater, logger *zap.Logger) UpdateTraining {
	return UpdateTraining{trainings: trainings, logger: logger}
}

// Update training plan godoc
//	@Summary		Updates a training plan
//	@Description	Updates a training plan with a name, description, difficulty, duration, exercises and tags
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training ID"
//	@Param			payload					body		trainings.UpdateTrainingRequest	true	"Body params"
//	@Success		200						{object}	models.TrainingPlan	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings{training_id}	[put]
func (h UpdateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.UpdateTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		validateErr := req.Validate()
		if err != nil || validateErr != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		req.ID = trainingID

		updatedTraining, err := h.trainings.UpdateTrainingPlan(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedTraining))
	}
}
