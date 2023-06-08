package middleware

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/gin-gonic/gin"
)

type TrainingID struct {
	TrainingID uint `uri:"trainingID" binding:"required"`
}

func BindTrainingIDFromUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var t TrainingID
		err := ctx.ShouldBindUri(&t)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			ctx.Abort()
			return
		}
		ctx.Set("trainingID", t.TrainingID)
	}
}
