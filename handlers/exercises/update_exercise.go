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

type UpdateExercise struct {
	exercises exercises.ExerciseUpdater
	logger    *zap.Logger
}

func NewUpdateExercise(exercises exercises.ExerciseUpdater, logger *zap.Logger) UpdateExercise {
	return UpdateExercise{exercises: exercises, logger: logger}
}

func (h UpdateExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req training.UpdateExerciseRequest
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
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrTrainingPlanNotFound))
				return
			}
			if errors.Is(err, contracts.ErrExerciseNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(contracts.ErrExerciseNotFound))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedTrainer) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(contracts.ErrUnauthorizedTrainer))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}
		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(updatedExercise))
	}
}
