package middleware

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/gin-gonic/gin"
)

type GoalID struct {
	GoalID uint `uri:"goalID" binding:"required"`
}

func BindGoalIDFromUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var g GoalID
		err := ctx.ShouldBindUri(&g)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			ctx.Abort()
			return
		}
		ctx.Set("goalID", g.GoalID)
	}
}
