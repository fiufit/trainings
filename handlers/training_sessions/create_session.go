package training_sessions

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/usecases/training_sessions"
	"github.com/gin-gonic/gin"
)

type CreateTrainingSession struct {
	uc training_sessions.TrainingSessionCreator
}

func NewCreateTrainingSession(uc training_sessions.TrainingSessionCreator) CreateTrainingSession {
	return CreateTrainingSession{uc: uc}
}

func (h CreateTrainingSession) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req tsContracts.CreateTrainingSessionRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		ts, err := h.uc.CreateTrainingSession(ctx, req.TrainingID, req.UserID)
		if err != nil {
			if errors.Is(err, contracts.ErrUserNotFound) || errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(ts))
	}
}
