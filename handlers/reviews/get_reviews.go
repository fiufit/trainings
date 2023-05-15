package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetReviews struct {
	logger *zap.Logger
}

func (h GetReviews) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
