package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateReview struct {
	logger *zap.Logger
}

func (h UpdateReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
