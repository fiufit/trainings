package server

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/fiufit/trainings/database"
	exerciseHandlers "github.com/fiufit/trainings/handlers/exercises"
	reviewHandlers "github.com/fiufit/trainings/handlers/reviews"
	trainingHandlers "github.com/fiufit/trainings/handlers/trainings"
	"github.com/fiufit/trainings/models"
	"github.com/fiufit/trainings/repositories"
	"github.com/fiufit/trainings/usecases/exercises"
	"github.com/fiufit/trainings/usecases/reviews"
	"github.com/fiufit/trainings/usecases/trainings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router         *gin.Engine
	createTraining trainingHandlers.CreateTraining
	getTrainings   trainingHandlers.GetTrainings
	updateTraining trainingHandlers.UpdateTraining
	deleteTraining trainingHandlers.DeleteTraining
	createExercise exerciseHandlers.CreateExercise
	deleteExercise exerciseHandlers.DeleteExercise
	updateExercise exerciseHandlers.UpdateExercise
	getExercise    exerciseHandlers.GetExercise
	createReview   reviewHandlers.CreateReview
	updateReview   reviewHandlers.UpdateReview
	getReviews     reviewHandlers.GetReviews
	getReviewByID  reviewHandlers.GetReviewByID
	deleteReview   reviewHandlers.DeleteReview
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

	err = db.AutoMigrate(&models.TrainingPlan{}, &models.Exercise{}, &models.Review{})
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
	reviewRepo := repositories.NewReviewRepository(db, logger)
	firebaseRepo, err := repositories.NewFirebaseRepository(logger, sdkJson, bucketName)
	if err != nil {
		panic(err)
	}

	// USECASES
	createTrainingUc := trainings.NewTrainingCreatorImpl(trainingRepo, userRepo, logger)
	getTrainingUc := trainings.NewTrainingGetterImpl(trainingRepo, firebaseRepo, logger)
	updateTrainingUc := trainings.NewTrainingUpdaterImpl(trainingRepo, firebaseRepo, logger)
	deleteTrainingUc := trainings.NewTrainingDeleterImpl(trainingRepo, logger)

	createExerciseUc := exercises.NewExerciseCreatorImpl(trainingRepo, exerciseRepo, logger)
	deleteExerciseUc := exercises.NewExerciseDeleterImpl(trainingRepo, exerciseRepo, logger)
	updateExerciseUc := exercises.NewExerciseUpdaterImpl(trainingRepo, exerciseRepo, logger)
	getExerciseUc := exercises.NewExerciseGetterImpl(trainingRepo, exerciseRepo, logger)

	createReviewUc := reviews.NewReviewCreatorImpl(trainingRepo, reviewRepo, userRepo, logger)
	getReviewUc := reviews.NewReviewGetterImpl(trainingRepo, reviewRepo, logger)
	updateReviewUc := reviews.NewReviewUpdaterImpl(trainingRepo, reviewRepo, logger)
	deleteReviewUc := reviews.NewReviewDeleterImpl(trainingRepo, reviewRepo, logger)

	// HANDLERS
	createTraining := trainingHandlers.NewCreateTraining(&createTrainingUc, logger)
	getTrainings := trainingHandlers.NewGetTrainings(&getTrainingUc, logger)
	updateTraining := trainingHandlers.NewUpdateTraining(&updateTrainingUc, logger)
	deleteTraining := trainingHandlers.NewDeleteTraining(&deleteTrainingUc, logger)

	createExercise := exerciseHandlers.NewCreateExercise(&createExerciseUc, logger)
	deleteExercise := exerciseHandlers.NewDeleteExercise(&deleteExerciseUc, logger)
	updateExercise := exerciseHandlers.NewUpdateExercise(&updateExerciseUc, logger)
	getExercise := exerciseHandlers.NewGetExercises(&getExerciseUc, logger)

	createReview := reviewHandlers.NewCreateReview(&createReviewUc, logger)
	getReviews := reviewHandlers.NewGetReviews(&getReviewUc, logger)
	getReviewByID := reviewHandlers.NewGetReviewByID(&getReviewUc, logger)
	updateReview := reviewHandlers.NewUpdateReview(&updateReviewUc, logger)
	deleteReview := reviewHandlers.NewDeleteReview(&deleteReviewUc, logger)

	return &Server{
		router:         gin.Default(),
		createTraining: createTraining,
		getTrainings:   getTrainings,
		updateTraining: updateTraining,
		createExercise: createExercise,
		deleteExercise: deleteExercise,
		updateExercise: updateExercise,
		getExercise:    getExercise,
		deleteTraining: deleteTraining,
		createReview:   createReview,
		getReviews:     getReviews,
		getReviewByID:  getReviewByID,
		updateReview:   updateReview,
		deleteReview:   deleteReview,
	}
}
