package handlers

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/usecases/exercises"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetExercise struct {
	exercises exercises.ExerciseGetter
	logger    *zap.Logger
}

func NewGetExercises(exercises exercises.ExerciseGetter, logger *zap.Logger) GetExercise {
	return GetExercise{exercises: exercises, logger: logger}
}

func (h GetExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req training.GetExerciseRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(string)
		exerciseID := ctx.MustGet("exerciseID").(string)
		req.TrainingPlanID = trainingID
		req.ExerciseID = exerciseID
		exercise, err := h.exercises.GetExerciseByID(ctx, req)
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
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(exercise))
	}
}
