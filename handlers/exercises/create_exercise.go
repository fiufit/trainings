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
