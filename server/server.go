package server

import (
	"fmt"
	"os"

	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/handlers"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"github.com/fiufit/trainings/usecases"
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
	db, err := database.NewPostgresDBClient()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.TrainingPlan{}, &models.Exercise{})
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewDevelopment()

	// REPOSITORIES
	trainingRepo := repositories.NewTrainingRepository(db, logger)

	// USECASES
	trainingUc := usecases.NewTrainingCreatorImpl(trainingRepo, logger)

	// HANDLERS
	createTraining := handlers.NewCreateTraining(&trainingUc, logger)
	getTrainings := handlers.NewGetTrainings(logger)

	return &Server{
		router:         gin.Default(),
		createTraining: createTraining,
		getTrainings:   getTrainings,
	}
}
