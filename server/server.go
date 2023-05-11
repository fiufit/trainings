package server

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/fiufit/trainings/database"
	"github.com/fiufit/trainings/handlers"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"github.com/fiufit/trainings/usecases/exercises"
	"github.com/fiufit/trainings/usecases/trainings"
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
	updateExercise handlers.UpdateExercise
	getExercise    handlers.GetExercise
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

	sdkJson, err := base64.StdEncoding.DecodeString(os.Getenv("FIREBASE_B64_SDK_JSON"))
	if err != nil {
		panic(err)
	}

	bucketName := os.Getenv("FIREBASE_BUCKET_NAME")

	// REPOSITORIES
	trainingRepo := repositories.NewTrainingRepository(db, logger)
	userRepo := repositories.NewUserRepository(usersUrl, logger, "v1")
	exerciseRepo := repositories.NewExerciseRepository(db, logger)
	firebaseRepo, err := repositories.NewFirebaseRepository(logger, sdkJson, bucketName)
	if err != nil {
		panic(err)
	}

	// USECASES
	createTrainingUc := trainings.NewTrainingCreatorImpl(trainingRepo, userRepo, logger)
	getTrainingUc := trainings.NewTrainingGetterImpl(trainingRepo, firebaseRepo, logger)
	updateTrainingUc := trainings.NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, logger)

	createExerciseUc := exercises.NewExerciseCreatorImpl(trainingRepo, exerciseRepo, logger)
	deleteExerciseUc := exercises.NewExerciseDeleterImpl(trainingRepo, exerciseRepo, logger)
	updateExerciseUc := exercises.NewExerciseUpdaterImpl(trainingRepo, exerciseRepo, logger)
	getExerciseUc := exercises.NewExerciseGetterImpl(trainingRepo, exerciseRepo, logger)

	// HANDLERS
	createTraining := handlers.NewCreateTraining(&createTrainingUc, logger)
	getTrainings := handlers.NewGetTrainings(&getTrainingUc, logger)
	updateTraining := handlers.NewUpdateTraining(&updateTrainingUc, logger)

	createExercise := handlers.NewCreateExercise(&createExerciseUc, logger)
	deleteExercise := handlers.NewDeleteExercise(&deleteExerciseUc, logger)
	updateExercise := handlers.NewUpdateExercise(&updateExerciseUc, logger)
	getExercise := handlers.NewGetExercises(&getExerciseUc, logger)

	return &Server{
		router:         gin.Default(),
		createTraining: createTraining,
		getTrainings:   getTrainings,
		updateTraining: updateTraining,
		createExercise: createExercise,
		deleteExercise: deleteExercise,
		updateExercise: updateExercise,
		getExercise:    getExercise,
	}
}
