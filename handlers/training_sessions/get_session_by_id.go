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
