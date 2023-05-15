package handlers

import (
	"errors"
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
			if errors.Is(err, contracts.ErrExerciseNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedTrainer) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(err))
				return
			}
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(""))
	}
}
