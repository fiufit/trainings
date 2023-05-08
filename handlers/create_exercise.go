package handlers

import (
	"errors"
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/fiufit/trainings/contracts/training"
	"github.com/fiufit/trainings/usecases"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateExercise struct {
	exercises usecases.ExerciseCreator
	logger    *zap.Logger
}

func NewCreateExercise(exercises usecases.ExerciseCreator, logger *zap.Logger) CreateExercise {
	return CreateExercise{exercises: exercises, logger: logger}
}

func (h CreateExercise) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req training.CreateExerciseRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}

		trainingID := ctx.MustGet("trainingID").(string)
		req.TrainingPlanID = trainingID

		res, err := h.exercises.CreateExercise(ctx, req)
		if err != nil {
			if errors.Is(err, contracts.ErrTrainingPlanNotFound) {
				ctx.JSON(http.StatusNotFound, contracts.FormatErrResponse(err))
				return
			}
			if errors.Is(err, contracts.ErrUnauthorizedTrainer) {
				ctx.JSON(http.StatusUnauthorized, contracts.FormatErrResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, contracts.FormatErrResponse(contracts.ErrInternal))
			return
		}

		ctx.JSON(http.StatusOK, contracts.FormatOkResponse(res))

	}
}
