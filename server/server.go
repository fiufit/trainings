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
	updateTraining handlers.UpdateTraining
	createExercise handlers.CreateExercise
	deleteExercise handlers.DeleteExercise
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
	usersUrl := os.Getenv("USERS_SERVICE_URL")

	// REPOSITORIES
	trainingRepo := repositories.NewTrainingRepository(db, logger)
	userRepo := repositories.NewUserRepository(usersUrl, logger, "v1")
	exerciseRepo := repositories.NewExerciseRepository(db, logger)

	// USECASES
	createTrainingUc := usecases.NewTrainingCreatorImpl(trainingRepo, userRepo, logger)
	getTrainingUc := usecases.NewTrainingGetterImpl(trainingRepo, logger)
	updateTrainingUc := usecases.NewTrainingUpdaterImpl(trainingRepo, logger)

	createExerciseUc := usecases.NewExerciseCreatorImpl(trainingRepo, exerciseRepo, logger)

	// HANDLERS
	createTraining := handlers.NewCreateTraining(&createTrainingUc, logger)
	getTrainings := handlers.NewGetTrainings(&getTrainingUc, logger)
	updateTraining := handlers.NewUpdateTraining(&updateTrainingUc, logger)

	createExercise := handlers.NewCreateExercise(&createExerciseUc, logger)

	return &Server{
		router:         gin.Default(),
		createTraining: createTraining,
		getTrainings:   getTrainings,
		updateTraining: updateTraining,
		createExercise: createExercise,
	}
}
