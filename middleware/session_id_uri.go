package middleware

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/gin-gonic/gin"
)

type TrainingSessionID struct {
	TrainingSessionID uint `uri:"trainingSessionID" binding:"required"`
}

func BindTrainingSessionIDFromUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ts TrainingSessionID
		err := ctx.ShouldBindUri(&ts)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		ctx.Set("trainingSessionID", ts.TrainingSessionID)
	}
}
