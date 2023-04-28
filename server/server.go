package server

import (
	"fmt"
	"os"

	"github.com/fiufit/trainings/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router         *gin.Engine
	createTraining handlers.CreateTraining
	getTrainings   handlers.GetTrainings
}

func (s *Server) Run() {
	err := s.router.Run(fmt.Sprintf("0.0.0.0:%v", os.Getenv("SERVICE_PORT")))
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	logger, _ := zap.NewDevelopment()

	// REPOSITORIES

	// USECASES

	// HANDLERS
	createTraining := handlers.NewCreateTraining(logger)
	getTrainings := handlers.NewGetTrainings(logger)

	return &Server{
		router:         gin.Default(),
		createTraining: createTraining,
		getTrainings:   getTrainings,
	}
}
