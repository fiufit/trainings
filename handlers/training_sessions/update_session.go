package training_sessions

import (
	"errors"
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

func (h UpdateTrainingSessions) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req tsContracts.UpdateTrainingSessionRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		ts, err := h.uc.UpdateTrainingSession(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrUnauthorizedAthlete) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(err))
				return
			}
			if errors.Is(err, contracts.ErrTrainingSessionAlreadyFinished) ||
				errors.Is(err, contracts.ErrTrainingSessionNotComplete) {

				ctx.JSON(http.StatusConflict, contracts.FormatErrResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(ts))
	}
}
