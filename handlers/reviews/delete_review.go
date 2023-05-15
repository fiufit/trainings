package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteReview struct {
	logger *zap.Logger
}

func (h DeleteReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
