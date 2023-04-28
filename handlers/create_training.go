package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateTraining struct {
	logger *zap.Logger
}

func NewCreateTraining(logger *zap.Logger) CreateTraining {
	return CreateTraining{logger: logger}
}

func (h CreateTraining) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotImplemented, "Not yet implemented")
	}
}
