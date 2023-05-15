package middleware

import (
	"net/http"

	"github.com/fiufit/trainings/contracts"
	"github.com/gin-gonic/gin"
)

type ReviewID struct {
	ReviewID uint `uri:"reviewID" binding:"required"`
}

func BindReviewIDFromUri() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r ReviewID
		err := ctx.ShouldBindUri(&r)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, contracts.FormatErrResponse(contracts.ErrBadRequest))
			return
		}
		ctx.Set("reviewID", r.ReviewID)
	}
}
