package training_sessions

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/usecases/training_sessions"
	"github.com/gin-gonic/gin"
)

type GetTrainingSessionByID struct {
	uc training_sessions.TrainingSessionGetter
}

func NewGetTrainingSessionByID(uc training_sessions.TrainingSessionGetter) GetTrainingSessionByID {
	return GetTrainingSessionByID{uc: uc}
}

//	Get training session by ID godoc
//	@Summary		Gets a training session by ID
//	@Description	Gets a training session by ID, validating that the requester has access to it
//	@Tags			training_sessions
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingSessionID		path		uint								true	"Training Session ID"
// 	@Param			requesterID					query		string							true	"Requester User ID"
//	@Success		200						{object}	models.TrainingSession	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/training_sessions{trainingSessionID}	[get]
func (h GetTrainingSessionByID) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trainingSessionID := ctx.MustGet("trainingSessionID").(uint)
		requesterID := ctx.Query("requester_id")

		ts, err := h.uc.GetByID(ctx, trainingSessionID, requesterID)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(ts))
	}
}
