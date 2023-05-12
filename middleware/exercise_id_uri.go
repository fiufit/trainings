package middleware

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/gin-gonic/gin"
)

type ExerciseID struct {
	ExerciseID uint `uri:"exerciseID" binding:"required"`
}

func BindExerciseIDFromUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var e ExerciseID
		err := ctx.ShouldBindUri(&e)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		ctx.Set("exerciseID", e.ExerciseID)
	}
}
