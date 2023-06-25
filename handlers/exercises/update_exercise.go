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
