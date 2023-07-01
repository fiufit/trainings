package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	uexercises "github.com/fiufit/trainings/usecases/exercises"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetExercise struct {
	exercises uexercises.ExerciseGetter
	logger    *zap.Logger
}

func NewGetExercises(exercises uexercises.ExerciseGetter, logger *zap.Logger) GetExercise {
	return GetExercise{exercises: exercises, logger: logger}
}

// Get exercise godoc
//	@Summary		Get exercise
//	@Description	Get an exercise for a given trainingID and exerciseID
//	@Tags			exercises
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint							true	"ID of the training to create the exercise for"
// 	@Param			exerciseID				path		uint							true	"ID of the exercise to delete"
//	@Success		200						{object}	models.Exercise	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/exercises/{exerciseID} 	[get]
func (h GetExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req exercises.GetExerciseRequest

		trainingID := ctx.MustGet("trainingID").(uint)
		exerciseID := ctx.MustGet("exerciseID").(uint)
		req.TrainingPlanID = trainingID
		req.ExerciseID = exerciseID
		exercise, err := h.exercises.GetExerciseByID(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(exercise))
	}
}
