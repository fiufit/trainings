package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateReview struct {
	//reviews reviews.ReviewCreator
	logger *zap.Logger
}

// func NewCreateReview(reviews reviews.ReviewCreator, logger *zap.Logger) CreateReview {
// 	return CreateReview{reviews: reviews, logger: logger}
// }

func (h CreateReview) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
