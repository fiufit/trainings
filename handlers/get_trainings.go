package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetTrainings struct {
	logger *zap.Logger
}

func NewGetTrainings(logger *zap.Logger) GetTrainings {
	return GetTrainings{logger: logger}
}

func (h GetTrainings) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotImplemented, "Not yet implemented")
	}
}
