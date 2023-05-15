package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetReviewByID struct {
	logger *zap.Logger
}

func (h GetReviewByID) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
