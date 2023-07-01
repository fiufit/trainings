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

//	Get training sessions godoc
//	@Summary		Get training sessions
//	@Description	Gets all training sessions of a user for a given training plan
//	@Tags			training_sessions
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			userID					query		string							true	"User ID"
// 	@Param			trainingID				query		uint								false	"Training Plan ID"
//	@Success		200						{object}	tsContracts.GetTrainingSessionsResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/training_sessions	[get]
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
