package training_sessions

import (
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

//	Create training session godoc
//	@Summary		Creates a new training session of a training plan
//	@Description	Creates a new review with a comment and a score for a given trainingPlanID and userID
//	@Tags			training_sessions
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			userID					query		string							true	"User ID"
// 	@Param			trainingID				query		uint								true	"Training Plan ID"
//	@Success		200						{object}	tsContracts.CreateTrainingSessionResponse	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/training_sessions	[post]
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
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(ts))
	}
}
