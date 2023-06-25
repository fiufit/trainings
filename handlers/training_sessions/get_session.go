package training_sessions

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	tsContracts "github.com/fiufit/trainings/contracts/training_sessions"
	"github.com/fiufit/trainings/usecases/training_sessions"
	"github.com/gin-gonic/gin"
)

type GetTrainingSessions struct {
	uc training_sessions.TrainingSessionGetter
}

func NewGetTrainingSessions(uc training_sessions.TrainingSessionGetter) GetTrainingSessions {
	return GetTrainingSessions{uc: uc}
}

func (h GetTrainingSessions) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req tsContracts.GetTrainingSessionsRequest
		err := ctx.ShouldBindQuery(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		req.Pagination.Validate()

		res, err := h.uc.Get(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))
	}
}
