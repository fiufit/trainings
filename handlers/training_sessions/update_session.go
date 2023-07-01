package training_sessions

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/usecases/training_sessions"
	"github.com/gin-gonic/gin"
)

type UpdateTrainingSessions struct {
	uc training_sessions.TrainingSessionUpdater
}

func NewUpdateTrainingSessions(uc training_sessions.TrainingSessionUpdater) UpdateTrainingSessions {
	return UpdateTrainingSessions{uc: uc}
}

//	Update training session godoc
//	@Summary		Updates a training session
//	@Description	Updates the progress of a training session, including the exercises and the step and minutes count
//	@Tags			training_sessions
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingSessionID		path		uint								true	"Training Session ID"
//	@Param			payload					body		tsContracts.UpdateTrainingSessionRequest true	"Body params"
//	@Success		200						{object}	tsContracts.UpdateTrainingSessionResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
// 	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
// 	@Failure		409						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/training_sessions{trainingSessionID}	[put]
func (h UpdateTrainingSessions) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req tsContracts.UpdateTrainingSessionRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		req.ID = ctx.MustGet("trainingSessionID").(uint)

		ts, err := h.uc.UpdateTrainingSession(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(ts))
	}
}
