package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	uexercises "github.com/fiufit/trainings/usecases/exercises"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateExercise struct {
	exercises uexercises.ExerciseCreator
	logger    *zap.Logger
}

func NewCreateExercise(exercises uexercises.ExerciseCreator, logger *zap.Logger) CreateExercise {
	return CreateExercise{exercises: exercises, logger: logger}
}

// Create exercise godoc
//	@Summary		Creates a new exercise
//	@Description	Creates a new exercise for a given trainingID, validating the training creator.
//	@Tags			exercises
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint							true	"ID of the training to create the exercise for"
//	@Param			payload					body		exercises.CreateExerciseRequest	true	"Body params"
//	@Success		200						{object}	models.Exercise	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/exercises 	[post]
func (h CreateExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req exercises.CreateExerciseRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		req.TrainingPlanID = trainingID

		res, err := h.exercises.CreateExercise(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))

	}
}
