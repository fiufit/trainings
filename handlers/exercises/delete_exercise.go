package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	uexercises "github.com/fiufit/trainings/usecases/exercises"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteExercise struct {
	exercises uexercises.ExerciseDeleter
	logger    *zap.Logger
}

func NewDeleteExercise(exercises uexercises.ExerciseDeleter, logger *zap.Logger) DeleteExercise {
	return DeleteExercise{exercises: exercises, logger: logger}
}

// Delete exercise godoc
//	@Summary		Delete an exercise
//	@Description	Delete an exercise for a given trainingID and exerciseID, validating the training creator.
//	@Tags			exercises
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint							true	"ID of the training to delete the exercise from"
// 	@Param			exerciseID				path		uint							true	"ID of the exercise to delete"
//	@Param			payload					body		exercises.DeleteExerciseRequest	true	"Body params"
//	@Success		200						{object}	string	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/exercises/{exerciseID} 	[delete]
func (h DeleteExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req exercises.DeleteExerciseRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		exerciseID := ctx.MustGet("exerciseID").(uint)
		req.TrainingPlanID = trainingID
		req.ExerciseID = exerciseID

		err = h.exercises.DeleteExercise(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
