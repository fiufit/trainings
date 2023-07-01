package handlers

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/exercises"
	uexercises "github.com/fiufit/trainings/usecases/exercises"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateExercise struct {
	exercises uexercises.ExerciseUpdater
	logger    *zap.Logger
}

func NewUpdateExercise(exercises uexercises.ExerciseUpdater, logger *zap.Logger) UpdateExercise {
	return UpdateExercise{exercises: exercises, logger: logger}
}

// Update exercise godoc
//	@Summary		Update exercise
//	@Description	Update an exercise for a given trainingID and exerciseID, validating the training creator
//	@Tags			exercises
//	@Accept			json
//	@Produce		json
//	@Param			version					path		string							true	"API Version"
// 	@Param			trainingID				path		uint							true	"ID of the training to update the exercise from"
// 	@Param			exerciseID				path		uint							true	"ID of the exercise to delete"
//	@Param			payload					body		exercises.UpdateExerciseRequest	true	"Body params"
//	@Success		200						{object}	models.Exercise	"Important Note: OK responses are wrapped in {"data": ... }"
//	@Failure		400						{object}	contracts.ErrResponse
//	@Failure		401						{object}	contracts.ErrResponse
//	@Failure		404						{object}	contracts.ErrResponse
//	@Failure		500						{object}	contracts.ErrResponse
//	@Router			/{version}/trainings/{trainingID}/exercises/{exerciseID} 	[patch]
func (h UpdateExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req exercises.UpdateExerciseRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(uint)
		exerciseID := ctx.MustGet("exerciseID").(uint)
		req.TrainingPlanID = trainingID
		req.ExerciseID = exerciseID

		updatedExercise, err := h.exercises.UpdateExercise(ctx, req)
		if err != nil {
			contracts.HandleErrorType(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedExercise))
	}
}
