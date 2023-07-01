package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/trainings"
	utrainings "github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteTraining struct {
	trainings utrainings.TrainingDeleter
	logger    *zap.Logger
}

func NewDeleteTraining(trainings utrainings.TrainingDeleter, logger *zap.Logger) DeleteTraining {
	return DeleteTraining{trainings: trainings, logger: logger}
}

// Delete training plan godoc
//	@Summary		Deletes a training plan
//	@Description	Deletes a training plan for a given trainingID and validates that the user has permissions to do so
//	@Tags			training_plans
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint								true	"Training ID"
//	@Param			payload					body		trainings.DeleteTrainingRequest	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings{training_id}	[delete]
func (h DeleteTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req trainings.DeleteTrainingRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		req.TrainingPlanID = trainingID

		err = h.trainings.DeleteTraining(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
