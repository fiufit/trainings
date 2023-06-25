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
